package client

import (
	"bytes"
	"context"
	"encoding/hex"
	"errors"
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"golang.org/x/net/publicsuffix"
)

type Client struct {
	endpointURL *url.URL
	secure      bool
	httpClient  *http.Client
	random      *rand.Rand // random seed
}

// Options for New method
type Options struct {
	Secure    bool
	Transport http.RoundTripper
}

// requestMetadata is container for all the values to make a request
type RequestMetadata struct {
	QueryValues   url.Values
	Authorization string
	ContentType   string
	ContentBody   io.Reader
	ContentLength int
	Headers       map[string]string
}

// lockedRandSource provides protected rand source, implements rand.Source interface
type LockedRandSource struct {
	lk  sync.Mutex
	src rand.Source
}

// list of success status
var successStatus = []int{
	http.StatusOK,
	http.StatusNoContent,
	http.StatusPartialContent,
}

func New(endpoint string, opts *Options) (*Client, error) {
	var (
		client *Client
		err    error
	)

	// construct endpoint
	endpointURL, err := getEndpointURL(endpoint, opts.Secure)
	if err != nil {
		return nil, err
	}

	// Initialize cookies to preserve server sent cookies if any and replay
	// them upon each request.
	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		return nil, err
	}

	transport := opts.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	// instantiate new Client
	client = &Client{
		endpointURL: endpointURL,
		secure:      opts.Secure,
		random:      rand.New(&LockedRandSource{src: rand.NewSource(time.Now().UTC().UnixNano())}),
		httpClient: &http.Client{
			Jar:       jar,
			Transport: transport,
		},
	}

	return client, err
}

// Int63 returns a non-negative pseudo-random 63-bit integer as an int64.
func (r *LockedRandSource) Int63() (n int64) {
	r.lk.Lock()
	n = r.src.Int63()
	r.lk.Unlock()
	return
}

// Seed uses the provided seed value to initialize the generator to a
// deterministic state.
func (r *LockedRandSource) Seed(seed int64) {
	r.lk.Lock()
	r.src.Seed(seed)
	r.lk.Unlock()
}

// executeMethod instantiates a given method, and retries the
// request upon any error up to maxRetries attempts in a binomially
// delayed manner using a standard back off algorithm
func (client *Client) ExecuteMethod(ctx context.Context, method string, path string, metadata RequestMetadata) (res *http.Response, err error) {
	var retryable bool       // Indicates if request can be retried.
	var bodySeeker io.Seeker // Extracted seeker from io.Reader.
	var reqRetry = MaxRetry  // Indicates how many times we can retry the request

	if strings.HasPrefix(path, "/") {
		return nil, errors.New("Path cannot have leading slash")
	}

	if metadata.ContentBody != nil {
		// Check if body is seekable then it is retryable.
		bodySeeker, retryable = metadata.ContentBody.(io.Seeker)
		switch bodySeeker {
		case os.Stdin, os.Stdout, os.Stderr:
			retryable = false
		}
		// retry only when reader is seekable
		if !retryable {
			reqRetry = 1
		}

		// Figure out if the body can be closed - if yes
		// we will definitely close it upon the function
		// return.
		bodyCloser, ok := metadata.ContentBody.(io.Closer)
		if ok {
			defer bodyCloser.Close()
		}
	}

	// Create cancel context to control 'newRetryTimer' go routine.
	retryCtx, cancel := context.WithCancel(ctx)
	// Indicate to our routine to exit cleanly upon return.
	defer cancel()

	for range client.newRetryTimer(retryCtx, reqRetry, DefaultRetryUnit, DefaultRetryCap, MaxJitter) {
		// Retry executes the following function body if request has an
		// error until maxRetries have been exhausted, retry attempts are
		// performed after waiting for a given period of time in a
		// binomial fashion
		if retryable {
			// seek back to beginning for each attempt
			if _, err = bodySeeker.Seek(0, 0); err != nil {
				// if seek failed, no need to retry
				return nil, err
			}
		}

		// Instantiate a new request
		var req *http.Request
		req, err = client.newRequest(ctx, method, path, metadata)
		if err != nil {
			continue
		}

		if metadata.Headers != nil {
			for headerName, headerValue := range metadata.Headers {
				req.Header.Set(headerName, headerValue)
			}
		}

		if metadata.Authorization != "" {
			req.Header.Set("Authorization", metadata.Authorization)
		}

		if metadata.ContentType != "" {
			req.Header.Set("Content-Type", metadata.ContentType)
		}

		// initiate the request
		res, err = client.do(req)
		if err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return nil, err
			}

			// retry the request
			continue
		}

		// for any known successful http status, return immediately
		for _, httpStatus := range successStatus {
			if httpStatus == res.StatusCode {
				return res, nil
			}
		}
	}

	// Return an error when retry is canceled or deadlined
	if e := retryCtx.Err(); e != nil {
		return nil, e
	}

	return res, err
}

// do executes http request
func (client *Client) do(req *http.Request) (*http.Response, error) {
	resp, err := client.httpClient.Do(req)
	if err != nil {
		// Handle this specifically for now until future Golang versions fix this issue properly.
		if urlErr, ok := err.(*url.Error); ok {
			if strings.Contains(urlErr.Err.Error(), "EOF") {
				return nil, &url.Error{
					Op:  urlErr.Op,
					URL: urlErr.URL,
					Err: errors.New("Connection closed by foreign host " + urlErr.URL + ". Retry again."),
				}
			}
		}
		return nil, err
	}

	// Response cannot be non-nil, report error if thats the case
	if resp == nil {
		return nil, errors.New("Response is empty")
	}

	return resp, nil
}

// newRequest - instantiate a new HTTP request for a given method.
func (client *Client) newRequest(ctx context.Context, method string, path string, metadata RequestMetadata) (req *http.Request, err error) {
	// If no method is supplied default to 'POST'.
	if method == "" {
		method = http.MethodPost
	}

	// Construct a new target URL.
	targetURL, err := client.makeTargetURL(
		path, metadata.QueryValues)
	if err != nil {
		return nil, err
	}

	// Initialize a new HTTP request for the method.
	req, err = http.NewRequestWithContext(ctx, method, targetURL.String(), nil)
	if err != nil {
		return nil, err
	}

	// Go net/http notoriously closes the request body.
	// - The request Body, if non-nil, will be closed by the underlying Transport, even on errors.
	// This can cause underlying *os.File seekers to fail, avoid that
	// by making sure to wrap the closer as a nop.
	if metadata.ContentLength == 0 {
		req.Body = nil
	} else {
		req.Body = ioutil.NopCloser(metadata.ContentBody)
	}

	// Set incoming content-length.
	req.ContentLength = int64(metadata.ContentLength)
	if req.ContentLength <= -1 {
		// For unknown content length, we upload using transfer-encoding: chunked.
		req.TransferEncoding = []string{"chunked"}
	}

	// Return request.
	return req, nil
}

func (client *Client) makeTargetURL(path string, queryValues url.Values) (*url.URL, error) {
	host := client.endpointURL.Host
	scheme := client.endpointURL.Scheme

	// strip port 80 and 443 so we won't send these ports in Host header.
	// The reason is that browsers and curl automatically remove :80 and :443
	// with the generated presigned urls, then a signature mismatch error.
	if h, p, err := net.SplitHostPort(host); err == nil {
		if scheme == "http" && p == "80" || scheme == "https" && p == "443" {
			host = h
		}
	}

	urlStr := scheme + "://" + host + "/" + path

	// If there are any query values, add them to the end.
	if len(queryValues) > 0 {
		urlStr = urlStr + "?" + queryEncode(queryValues)
	}

	return url.Parse(urlStr)
}

// QueryEncode encodes query values in their URL encoded form. In
// addition to the percent encoding performed by urlEncodePath() used
// here, it also percent encodes '/' (forward slash)
func queryEncode(v url.Values) string {
	if v == nil {
		return ""
	}
	var buf bytes.Buffer
	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		vs := v[k]
		prefix := percentEncodeSlash(encodePath(k)) + "="
		for _, v := range vs {
			if buf.Len() > 0 {
				buf.WriteByte('&')
			}
			buf.WriteString(prefix)
			buf.WriteString(percentEncodeSlash(encodePath(v)))
		}
	}
	return buf.String()
}

// if object matches reserved string, no need to encode them
var reservedObjectNames = regexp.MustCompile("^[a-zA-Z0-9-_.~/]+$")

// EncodePath encode the strings from UTF-8 byte representations to HTML hex escape sequences
//
// This is necessary since regular url.Parse() and url.Encode() functions do not support UTF-8
// non english characters cannot be parsed due to the nature in which url.Encode() is written
//
// This function on the other hand is a direct replacement for url.Encode() technique to support
// pretty much every UTF-8 character.
func encodePath(pathName string) string {
	if reservedObjectNames.MatchString(pathName) {
		return pathName
	}
	var encodedPathname strings.Builder
	for _, s := range pathName {
		if 'A' <= s && s <= 'Z' || 'a' <= s && s <= 'z' || '0' <= s && s <= '9' { // §2.3 Unreserved characters (mark)
			encodedPathname.WriteRune(s)
			continue
		}
		switch s {
		case '-', '_', '.', '~', '/': // §2.3 Unreserved characters (mark)
			encodedPathname.WriteRune(s)
			continue
		default:
			len := utf8.RuneLen(s)
			if len < 0 {
				// if utf8 cannot convert return the same string as is
				return pathName
			}
			u := make([]byte, len)
			utf8.EncodeRune(u, s)
			for _, r := range u {
				hex := hex.EncodeToString([]byte{r})
				encodedPathname.WriteString("%" + strings.ToUpper(hex))
			}
		}
	}
	return encodedPathname.String()
}

// Expects ascii encoded strings - from output of urlEncodePath
func percentEncodeSlash(s string) string {
	return strings.Replace(s, "/", "%2F", -1)
}
