package client

import (
	"context"
	"io"
	"testing"
)

func TestClient(t *testing.T) {
	instant, err := New("www.google.com", &Options{
		Secure: true,
	})

	if err != nil {
		t.Error(err.Error())
		return
	}

	ctx := context.Background()

	res, err := instant.ExecuteMethod(ctx, "POST", "", RequestMetadata{
		// ContentType:   "application/json",
		// ContentLength: len(byteSlice),
		// ContentBody:   bytes.NewReader(byteSlice),
	})

	if err != nil {
		t.Error(err.Error())
		return
	}

	t.Logf("response status: %v\n", res.StatusCode)
	if b, err := io.ReadAll(res.Body); err == nil {
		t.Logf("response body: %v\n", string(b))
	} else {
		t.Errorf("no response body can be displayed\n")
	}
}
