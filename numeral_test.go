package main

import (
	"testing"
)

func TestNumeral(t *testing.T) {
	t.Run("1 converts to I", func(t *testing.T) {
		got := ConvertToRoman(1)
		want := "I"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("2 converts to II", func(t *testing.T) {
		got := ConvertToRoman(2)
		want := "II"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
