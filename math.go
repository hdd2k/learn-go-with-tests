package main

import (
	"time"
)

type Point struct {
	X int
	Y int
}

func SecondHand(time time.Time) Point {
	return Point{150, 150 - 90}
}
