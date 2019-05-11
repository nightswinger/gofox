package sigfox

import (
	"context"
	"testing"
)

func TestNewClient(t *testing.T) {
	c, _ := NewClient("LOGIN_ID", "PASSWORD")

	if got, want := c.baseURL.String(), defaultBaseURL; got != want {
		t.Errorf("NewClient baseURL is %v, want %v", got, want)
	}
}

func TestNewRequest(t *testing.T) {
	c, _ := NewClient("LOGIN_ID", "PASSWORD")
	ctx := context.Background()

	inURL, outURL := "/foo", defaultBaseURL+"/foo"
	req, _ := c.newRequest(ctx, "GET", inURL, nil)

	// test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("newRequest(%q) URL is %v, want %v", inURL, got, want)
	}
}
