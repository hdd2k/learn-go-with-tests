package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Hank")

	got := buffer.String()
	want := "Hello, Hank"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
