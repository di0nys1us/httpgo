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
	w.Header().Set(HeaderContentType, MediaTypeJSON)
	w.WriteHeader(statusCode)

	if body != nil {
		return json.NewEncoder(w).Encode(body)
	}

	return nil
}

type Response struct {
	StatusCode int
	Headers    map[string]string
	Cookies    []*http.Cookie
	Body       interface{}
}

func NewResponse(statusCode int) *Response {
	return &Response{statusCode, map[string]string{}, nil, nil}
}

func (r *Response) WithHeader(key, value string) *Response {
	headers := map[string]string{key: value}

	for k, v := range r.Headers {
		headers[k] = v
	}

	return &Response{r.StatusCode, headers, r.Cookies, r.Body}
}

func (r *Response) WithCookie(c *http.Cookie) *Response {
	return &Response{r.StatusCode, r.Headers, append(r.Cookies, c), r.Body}
}

func (r *Response) WithBody(body interface{}) *Response {
	return &Response{r.StatusCode, r.Headers, r.Cookies, body}
}

func ResponseOK() *Response {
	return NewResponse(http.StatusOK)
}

func ResponseCreated() *Response {
	return NewResponse(http.StatusCreated)
}

func ResponseNoContent() *Response {
	return NewResponse(http.StatusNoContent)
}

func ResponseBadRequest() *Response {
	return NewResponse(http.StatusBadRequest)
}

func ResponseUnauthorized() *Response {
	return NewResponse(http.StatusUnauthorized)
}

func ResponseForbidden() *Response {
	return NewResponse(http.StatusForbidden)
}

func ResponseNotFound() *Response {
	return NewResponse(http.StatusNotFound)
}

func ResponseInternalServerError() *Response {
	return NewResponse(http.StatusInternalServerError)
}
