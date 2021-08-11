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

	// checkArea := func(t testing.TB, shape Shape, expected float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	if got != expected {
	// 		t.Errorf("got %.2g expected %.2g", got, expected)
	// 	}
	// }

	areaTests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expected: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, expected: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, expected: 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.expected {
			t.Errorf("got %g expected %g", got, tt.expected)
		}
	}

	// t.Run("rectangle area", func(t *testing.T) {
	// 	rectangle := Rectangle{10.0, 10.0}
	// 	expected := 100.0
	// 	checkArea(t, rectangle, expected)
	// })

	// t.Run("circle area", func(t *testing.T) {
	// 	circle := Circle{10.0}
	// 	expected := 314.1592653589793
	// 	checkArea(t, circle, expected)
	// })

}
