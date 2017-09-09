package httpgo

import "testing"
import "net/http"
import "net/http/httptest"

const (
	location = "http://localhost/resource/1"
)

func createResponseHandlerFunc(response *Response) ResponseHandlerFunc {
	return ResponseHandlerFunc(func(w http.ResponseWriter, r *http.Request) (*Response, error) {
		return response, nil
	})
}

func TestServeHTTP(t *testing.T) {
	handler := createResponseHandlerFunc(ResponseOK(true))

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	handler.ServeHTTP(w, r)

	if statusCode := w.Result().StatusCode; statusCode != http.StatusOK {
		t.Errorf("got %v, want %v", statusCode, http.StatusOK)
	}

	handler = createResponseHandlerFunc(ResponseCreated(true, location))

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/", nil)

	handler.ServeHTTP(w, r)

	if statusCode := w.Result().StatusCode; statusCode != http.StatusCreated {
		t.Errorf("got %v, want %v", statusCode, http.StatusCreated)
	}

	if locationHeader := w.Result().Header.Get(HeaderLocation); locationHeader != location {
		t.Errorf("got %v, want %v", 0, 0)
	}
}
