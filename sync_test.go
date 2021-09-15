package main

import (
	"testing"
)

// TestCounter tests a concurrency-safe counter
func TestCounter(t *testing.T) {
	// start with not thread-safe counter
	t.Run("Increment counter 3 times to get 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, expected %d", got.Value(), 3)
	}
}
