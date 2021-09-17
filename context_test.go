package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	expected := "hello, world"
	server := Server(&StubStore{expected})
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res := httptest.NewRecorder()

	server.ServeHTTP(res, req)

	if res.Body.String() != expected {
		t.Errorf("got '%s' want '%s'", res.Body.String(), expected)
	}

}
