package httpgo

import (
	"net/http"
)

func ErrorWithStatusCode(statusCode int, cause error) *withStatusCode {
	return &withStatusCode{statusCode, cause}
}

type withStatusCode struct {
	statusCode int
	cause      error
}

func (err *withStatusCode) Error() string {
	return http.StatusText(err.statusCode)
}

func (err *withStatusCode) Cause() error {
	return err.cause
}

func (err *withStatusCode) StatusCode() int {
	return err.statusCode
}

func (err *withStatusCode) Body() interface{} {
	return nil
}
