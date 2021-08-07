package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()
	expected := 40.0
	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, expected float64) {
		t.Helper()
		got := shape.Area()
		if got != expected {
			t.Errorf("got %.2g expected %.2g", got, expected)
		}
	}

	t.Run("rectangle area", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		expected := 100.0
		checkArea(t, rectangle, expected)
	})

	t.Run("circle area", func(t *testing.T) {
		circle := Circle{10.0}
		expected := 314.1592653589793
		checkArea(t, circle, expected)
	})

}
