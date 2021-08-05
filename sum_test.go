package main

import (
	"reflect"
	"testing"
)

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

func TestSumAll(t *testing.T) {
	sumAlled := SumAll(
		[]int{1, 2, 4},
		[]int{4, 5, 7},
	)
	expected := []int{7, 16}
	if !reflect.DeepEqual(sumAlled, expected) {
		t.Errorf("expected %v got %v", expected, sumAlled)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(tb testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("expected %v got %v", want, got)
		}
	}

	t.Run("SumAllTail with 1+ len arrs", func(t *testing.T) {
		sumAllTailed := SumAllTail(
			[]int{1, 2, 4},
			[]int{4, 5, 7},
		)
		expected := []int{6, 12}
		checkSums(t, sumAllTailed, expected)
	})
	t.Run("SumAllTail with 0 len arrs", func(t *testing.T) {
		sumAllTailed := SumAllTail(
			[]int{},
			[]int{},
		)
		expected := []int{0, 0}
		checkSums(t, sumAllTailed, expected)
	})
}
