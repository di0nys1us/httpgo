package httpgo

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

const (
	locationTest = "http://localhost/resource/1"
)

type errorTest string

func (r errorTest) Error() string {
	return string(r)
}

type errorStatusCode int

func (c errorStatusCode) StatusCode() int {
	return int(c)
}

type errorBody string

func (b errorBody) Body() interface{} {
	return string(b)
}

func TestResponseHandlerFunc(t *testing.T) {
	t.Run("TestServeHTTP", func(t *testing.T) {
		cases := []struct {
			name       string
			handler    ErrorHandlerFunc
			statusCode int
			body       interface{}
		}{
			{
				"return nil",
				ErrorHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
					return nil
				}),
				http.StatusOK,
				nil,
			},
			{
				"return error",
				ErrorHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
					return errors.New("test")
				}),
				http.StatusInternalServerError,
				`{ "message": "test" }`,
			},
			{
				"return custom error",
				ErrorHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
					return struct {
						errorTest
					}{"test"}
				}),
				http.StatusInternalServerError,
				`{ "message": "test" }`,
			},
			{
				"return error with custom status code",
				ErrorHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
					return struct {
						errorTest
						errorStatusCode
					}{"test", http.StatusBadRequest}
				}),
				http.StatusBadRequest,
				`{ "message": "test" }`,
			},
			{
				"return error with custom body",
				ErrorHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
					return struct {
						errorTest
						errorBody
					}{"test", "true"}
				}),
				http.StatusInternalServerError,
				true,
			},
			{
				"return error with custom status code and body",
				ErrorHandlerFunc(func(w http.ResponseWriter, r *http.Request) error {
					return struct {
						errorTest
						errorStatusCode
						errorBody
					}{"test", http.StatusBadRequest, "true"}
				}),
				http.StatusBadRequest,
				true,
			},
		}

		for _, c := range cases {
			t.Run(c.name, func(t *testing.T) {
				w := httptest.NewRecorder()
				r := httptest.NewRequest(http.MethodGet, "/", nil)

				c.handler.ServeHTTP(w, r)

				assert.Exactly(t, c.statusCode, w.Code)

				if body, ok := c.body.(string); ok {
					assert.JSONEq(t, body, w.Body.String())
				}
			})
		}
	})
}
