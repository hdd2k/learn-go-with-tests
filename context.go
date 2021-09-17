package main

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

type Store interface {
	Fetch() string
	Cancel()
}

// Server creation
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case d := <-data:
			fmt.Fprint(w, d)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}

// Store spy
type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Error("store was NOT cancelled")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Error("store was cancelled")
	}
}
