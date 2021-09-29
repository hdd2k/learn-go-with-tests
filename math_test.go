package main

import (
	"math"
	"testing"
	"time"
)

// every clock has a centre of (150, 150)
// the hour hand is 50 long
// the minute hand is 80 long
// the second hand is 90 long.

func TestSecondHandAtMidnight(t *testing.T) {
	time := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := Point{X: 150, Y: 150 - 90}
	got := SecondHand(time)

	if want != got {
		t.Errorf("Got %v want %v", got, want)
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {
	time := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

	want := Point{X: 150, Y: 150 + 90}
	got := SecondHand(time)

	if want != got {
		t.Errorf("Got %v want %v", got, want)
	}
}

// write smaller tests first - comment out integration test during this
func TestSecondsToRadians(t *testing.T) {
	tests := []struct {
		time time.Time
		rad  float64
	}{
		// simpleTime(hour, min, second)
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, test := range tests {
		t.Run(testName(test.time), func(t *testing.T) {
			want := test.rad
			got := secondsInRadians(test.time)
			if want != got {
				t.Errorf("Want %v radians got %v", want, got)
			}
		})
	}
}

func TestSecondsToVector(t *testing.T) {
	tests := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{0, -1}},
		{simpleTime(0, 0, 0), Point{0, 1}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, test := range tests {
		t.Run(testName(test.time), func(t *testing.T) {
			want := test.point
			got := secondHandPoint(test.time)
			if !roughlyEqualPoint(want, got) {
				t.Errorf("Want %v Point got %v Point", want, got)
			}
		})
	}

}

func simpleTime(hour, min, second int) time.Time {
	return time.Date(1337, time.January, 1, hour, min, second, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}

func roughlyEqualPoint(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func roughlyEqualFloat64(a, b float64) bool {
	const threshold = 1e-7
	return math.Abs(a-b) < threshold
}
