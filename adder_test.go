package main

import (
	"fmt"
	"testing"
)

// Test is run if prefixed with "Test"
func TestAddr(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// Example is run if prefixed with "Example" --- and the "Output" comment at the end causes example run (without it no Example Run)
func ExampleAdd() {
	sum := Add(2, 5)
	fmt.Println(sum)
}
