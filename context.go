package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server creation
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		res, err := store.Fetch(r.Context())

		if err != nil {
			return
		}

		fmt.Fprint(w, res)
	}
}

// Store spy
type SpyStore struct {
	response string
	t        *testing.T
}

// spy response writer - need to check whether Response was written or not
type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("Not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {

	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store was cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case d := <-data:
		return d, nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
