package httpgo

import (
	"log"
	"net/http"
)

type errorMessage struct {
	Message string `json:"message"`
}

func newErrorMessage(err error) *errorMessage {
	return &errorMessage{err.Error()}
}

// ErrorHandlerFunc TODO
type ErrorHandlerFunc func(http.ResponseWriter, *http.Request) error

// ServeHTTP TODO
func (f ErrorHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := f(w, r)

	if err == nil {
		return
	}

	log.Printf("httpgo: %v", err)

	switch v := err.(type) {

	case ErrorResponse:
		WriteJSON(w, v.StatusCode(), v.Body())

	case StatusCodeResponse:
		WriteJSON(w, v.StatusCode(), newErrorMessage(err))

	case BodyResponse:
		WriteJSON(w, http.StatusInternalServerError, v.Body())

	default:
		WriteJSON(w, http.StatusInternalServerError, newErrorMessage(err))

	}
}
