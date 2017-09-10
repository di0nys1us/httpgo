package httpgo

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	locationTest = "http://localhost/resource/1"
)

var (
	errTest = errors.New("testMessage")
)

func createResponseHandlerFunc(response *Response, err error) ResponseHandlerFunc {
	return ResponseHandlerFunc(func(w http.ResponseWriter, r *http.Request) (*Response, error) {
		return response, err
	})
}

func TestServeHTTP(t *testing.T) {
	handler := createResponseHandlerFunc(ResponseOK().WithBody(true), nil)

	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)

	handler.ServeHTTP(w, r)

	if statusCode := w.Result().StatusCode; statusCode != http.StatusOK {
		t.Errorf("got %v, want %v", statusCode, http.StatusOK)
	}

	handler = createResponseHandlerFunc(ResponseCreated().WithBody(true).WithHeader(HeaderLocation, locationTest), nil)

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/", nil)

	handler.ServeHTTP(w, r)

	if statusCode := w.Result().StatusCode; statusCode != http.StatusCreated {
		t.Errorf("got %v, want %v", statusCode, http.StatusCreated)
	}

	if locationHeader := w.Result().Header.Get(HeaderLocation); locationHeader != locationTest {
		t.Errorf("got %v, want %v", locationHeader, locationTest)
	}

	handler = createResponseHandlerFunc(nil, errTest)

	w = httptest.NewRecorder()
	r = httptest.NewRequest(http.MethodGet, "/", nil)

	handler.ServeHTTP(w, r)

	if statusCode := w.Result().StatusCode; statusCode != http.StatusInternalServerError {
		t.Errorf("got %v, want %v", statusCode, http.StatusInternalServerError)
	}

	errorMessage := &errorMessage{}
	err := ReadJSON(w.Result().Body, errorMessage)

	if err != nil {
		t.Error("got err, want nil")
	}

	if errorMessage.Message != errTest.Error() {
		t.Errorf("got %v, want %v", errorMessage.Message, errTest.Error())
	}
}
