package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	t.Run("basic happy path request", func(t *testing.T) {
		expected := "hello, world"
		store := &SpyStore{response: expected, t: t}
		server := Server(store)
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		if res.Body.String() != expected {
			t.Errorf("got '%s' want '%s'", res.Body.String(), expected)
		}
	})
	t.Run("tell store to cancel work if request cancelled", func(t *testing.T) {
		expected := "hello, world"
		store := &SpyStore{response: expected, t: t}
		server := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(cancellingCtx)

		res := &SpyResponseWriter{}

		server.ServeHTTP(res, req)

		if res.written {
			t.Error("response was written, but should NOT have")
		}
	})

}
