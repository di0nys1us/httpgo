package httpgo

import (
	"net/http"
)

type errorMessage struct {
	Message string `json:"message"`
}

type ResponseHandlerFunc func(http.ResponseWriter, *http.Request) (*Response, error)

func (f ResponseHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response, err := f(w, r)

	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, &errorMessage{err.Error()})
	}

	if err == nil && response != nil {
		for k, v := range response.Headers {
			w.Header().Set(k, v)
		}

		for _, c := range response.Cookies {
			http.SetCookie(w, &c)
		}

		WriteJSON(w, response.StatusCode, response.Body)
	}
}
