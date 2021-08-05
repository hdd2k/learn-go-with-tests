package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("collection of any nums", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		summed := Sum(numbers)
		expected := 6
		if summed != expected {
			t.Errorf("expected %d but got %d, given %v", expected, summed, numbers)
		}
	})
}
