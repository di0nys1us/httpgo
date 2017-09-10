package httpgo

import (
	"net/http"
	"testing"
)

func TestResponse(t *testing.T) {
	r := NewResponse(http.StatusOK).
		WithHeader("foo", "bar").
		WithHeader("bar", "foo")

	if c := r.StatusCode; c != http.StatusOK {
		t.Errorf("got %v, want %v", c, http.StatusOK)
	}

	if h := r.Headers["foo"]; h != "bar" {
		t.Errorf("got %v, want %v", h, "bar")
	}

	if h := r.Headers["bar"]; h != "foo" {
		t.Errorf("got %v, want %v", h, "foo")
	}
}
