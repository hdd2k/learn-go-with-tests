package main

import (
	"sync"
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

	t.Run("Thread-safe increment test", func(t *testing.T) {
		wantCount := 1000
		counter := Counter{}

		// waitgroup - convenient way to synchronize concurrent processes/threads
		var wg sync.WaitGroup
		// Add() -> set the NUMBER of processes to WAIT for
		wg.Add(wantCount)

		for i := 0; i < wantCount; i++ {
			go func() {
				counter.Inc()
				// Registers 1 "Done" process to the waitgroup
				wg.Done()
			}()
		}
		// waits until ALL expected processes are finished
		wg.Wait()

		assertCounter(t, &counter, wantCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, expected %d", got.Value(), want)
	}
}
