package httpgo

import (
	"net/http"
)

type ResponseHandlerFunc func(http.ResponseWriter, *http.Request) (*Response, error)

func (f ResponseHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response, err := f(w, r)

	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, &errorMessage{err.Error()})
		return
	}

	if response != nil {
		for k, v := range response.Headers {
			w.Header().Set(k, v)
		}

		WriteJSON(w, response.StatusCode, response.Body)
	}
}
