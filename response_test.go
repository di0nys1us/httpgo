package httpgo

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var testJSON = `{ "id": 1000, "name": "john" }`

func TestReadJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s := &testStruct{}

		err := ReadJSON(strings.NewReader(testJSON), s)

		assert.Nil(t, err)
		assert.Exactly(t, &testStruct{1000, "john"}, s)
	})

	t.Run("failure", func(t *testing.T) {
		s := &testStruct{}

		err := ReadJSON(strings.NewReader(`{ "test" }`), s)

		assert.Error(t, err)
	})
}

func TestWriteJSON(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		w := httptest.NewRecorder()

		err := WriteJSON(w, http.StatusAccepted, &testStruct{1000, "john"})

		assert.Nil(t, err)
		assert.Exactly(t, http.StatusAccepted, w.Code)
		assert.JSONEq(t, testJSON, w.Body.String())
	})

	t.Run("body is nil", func(t *testing.T) {
		w := httptest.NewRecorder()

		err := WriteJSON(w, http.StatusAccepted, nil)

		assert.Nil(t, err)
		assert.Exactly(t, http.StatusAccepted, w.Code)
		assert.Exactly(t, "", w.Body.String())
	})
}
