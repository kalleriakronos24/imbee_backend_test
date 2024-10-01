package api_test

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestGetMessage(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		msg := `{"message": "hello"}`
		_, _ = w.Write([]byte(msg))
		w.WriteHeader(http.StatusOK)
	}

	apitest.New(). // configuration
			HandlerFunc(handler).
			Get("/message"). // request
			Expect(t).       // expectations
			Body(`{"message": "hello"}`).
			Status(http.StatusOK).
			End()
}
