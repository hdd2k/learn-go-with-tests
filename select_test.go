package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got, _ := Race(BaseRacer, slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

	t.Run("timeout if longer than 10 sec", func(t *testing.T) {
		serverA := makeDelayedServer(time.Second * 11)
		serverB := makeDelayedServer(time.Second * 12)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Race(SelectRacer, serverA.URL, serverB.URL)

		if err == nil {
			t.Error("Expected timeout error but none found")
		}

	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
