package httpgo

import (
	"encoding/json"
	"io"
	"net/http"
)

// StatusCodeResponse TODO
type StatusCodeResponse interface {
	StatusCode() int
}

// BodyResponse TODO
type BodyResponse interface {
	Body() interface{}
}

// ErrorResponse TODO
type ErrorResponse interface {
	StatusCodeResponse
	BodyResponse
}

// ReadJSON TODO
func ReadJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

// WriteJSON TODO
func WriteJSON(w http.ResponseWriter, statusCode int, body interface{}) error {
	w.Header().Set(HeaderContentType, MediaTypeJSON)
	w.WriteHeader(statusCode)

	if body == nil {
		return nil
	}

	return json.NewEncoder(w).Encode(body)
}
