package main

import (
	"testing"
	"time"
)

// every clock has a centre of (150, 150)
// the hour hand is 50 long
// the minute hand is 80 long
// the second hand is 90 long.

func TestSecondHanAtMidnight(t *testing.T) {
	time := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	want := Point{X: 150, Y: 150 - 90}
	got := SecondHand(time)

	if want != got {
		t.Errorf("Got %v want %v", got, want)
	}
}
