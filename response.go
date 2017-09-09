package httpgo

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadJSON(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func WriteJSON(w http.ResponseWriter, statusCode int, body interface{}) error {
	w.Header().Set(HeaderContentType, MediaTypeApplicationJSONUTF8)
	w.WriteHeader(statusCode)

	if body != nil {
		return json.NewEncoder(w).Encode(body)
	}

	return nil
}

type errorMessage struct {
	Message string `json:"message"`
}

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       interface{}
	Error      error
}

func ResponseOK(body interface{}) *Response {
	return &Response{http.StatusOK, nil, body, nil}
}

func ResponseCreated(body interface{}, location string) *Response {
	return &Response{http.StatusCreated, map[string]string{HeaderLocation: location}, body, nil}
}

func ResponseNoContent() *Response {
	return &Response{http.StatusNoContent, nil, nil, nil}
}

func ResponseBadRequest() *Response {
	return &Response{http.StatusBadRequest, nil, nil, nil}
}

func ResponseUnauthorized() *Response {
	return &Response{http.StatusUnauthorized, nil, nil, nil}
}

func ResponseForbidden() *Response {
	return &Response{http.StatusForbidden, nil, nil, nil}
}

func ResponseNotFound() *Response {
	return &Response{http.StatusNotFound, nil, nil, nil}
}

func ResponseInternalServerError(err error) *Response {
	return &Response{http.StatusInternalServerError, nil, nil, err}
}
