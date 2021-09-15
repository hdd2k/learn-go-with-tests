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

		if counter.Value() != 3 {
			t.Errorf("got %d, expected %d", counter.Value(), 3)
		}
	})
}
