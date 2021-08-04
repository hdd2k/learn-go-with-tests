package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	// Test helper
	assertCorrectMsg := func(t testing.TB, got, expected string) {
		t.Helper() // marks this function as helper, and supplies _caller_ as error location:w
		if got != expected {
			t.Errorf("got %q want %q", got, expected)
		}
	}
	// subtests
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		expected := "Hello, Chris"
		assertCorrectMsg(t, got, expected)
	})
	t.Run("say 'Hello, World!' when empty string supplied", func(t *testing.T) {
		got := Hello("", "")
		expected := "Hello, World"
		assertCorrectMsg(t, got, expected)
	})
	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		expected := "Hola, Elodie"
		assertCorrectMsg(t, got, expected)
	})
	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		expected := "Bonjour, Elodie"
		assertCorrectMsg(t, got, expected)
	})
}
