package client

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
)

var (
	blankURL = url.URL{}
)

// isValidDomain validates if input string is a valid domain name
func isValidDomain(host string) bool {
	// See RFC 1035, RFC 3696
	host = strings.TrimSpace(host)
	if len(host) == 0 || len(host) > 255 {
		return false
	}
	// host cannot start or end with "-"
	if host[len(host)-1:] == "-" || host[:1] == "-" {
		return false
	}
	// host cannot start or end with "_"
	if host[len(host)-1:] == "_" || host[:1] == "_" {
		return false
	}
	// host cannot start or end with a "."
	if host[len(host)-1:] == "." || host[:1] == "." {
		return false
	}
	// All non alphanumeric characters are invalid.
	if strings.ContainsAny(host, "`~!@#$%^&*()+={}[]|\\\"';:><?/") {
		return false
	}
	// No need to regexp match, since the list is non-exhaustive
	// We let it valid and fail later
	return true
}

// isValidIP parses input string for ip address validity.
func isValidIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// Verify if input endpoint URL is valid.
func isValidEndpointURL(endpointURL url.URL) error {
	if endpointURL == blankURL {
		return errors.New("Endpoint url cannot be empty")
	}
	if endpointURL.Path != "/" && endpointURL.Path != "" {
		return errors.New("Endpoint url cannot have fully qualified paths")
	}

	return nil
}

// getEndpointURL construct a new endpoint
func getEndpointURL(endpoint string, secure bool) (*url.URL, error) {
	if strings.Contains(endpoint, ":") {
		host, _, err := net.SplitHostPort(endpoint)
		if err != nil {
			return nil, err
		}
		if !isValidIP(host) && !isValidDomain(host) {
			return nil, fmt.Errorf("'%s' does not follow ip address or domain name standards", endpoint)
		}
	} else {
		if !isValidIP(endpoint) && !isValidDomain(endpoint) {
			return nil, fmt.Errorf("'%s' does not follow ip address or domain name standards", endpoint)
		}
	}
	// If secure is false, use 'http' scheme
	scheme := "https"
	if !secure {
		scheme = "http"
	}

	// Construct a secured endpoint URL
	endpointURLStr := scheme + "://" + endpoint
	endpointURL, err := url.Parse(endpointURLStr)
	if err != nil {
		return nil, err
	}

	// Validate incoming endpoint URL
	if err := isValidEndpointURL(*endpointURL); err != nil {
		return nil, err
	}
	return endpointURL, nil
}

// CloseResponse close non nil response with any response Body.
// convenient wrapper to drain any remaining data on response body.
//
// Subsequently this allows golang http RoundTripper
// to re-use the same connection for future requests.
func CloseResponse(resp *http.Response) {
	// Callers should close resp.Body when done reading from it.
	// If resp.Body is not closed, the Client's underlying RoundTripper
	// (typically Transport) may not be able to re-use a persistent TCP
	// connection to the server for a subsequent "keep-alive" request.
	if resp != nil && resp.Body != nil {
		// Drain any remaining Body and then close the connection.
		// Without this closing connection would disallow re-using
		// the same connection for future uses.
		//  - http://stackoverflow.com/a/17961593/4465767
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}
}

func ResetQuery(m map[string][]string) map[string][]string {
	dicts := make(map[string][]string)
	for k, v := range m {
		lists := strings.Split(k, "&")
		if len(lists) == 1 {
			dicts[k] = v
			continue
		}
		for _, vv := range lists {
			p := strings.Split(vv, "=")
			dicts[p[0]] = append(dicts[p[0]], p[1])
		}
	}
	return dicts
}
