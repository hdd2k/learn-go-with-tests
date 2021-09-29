package main

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(time time.Time) Point {
	// angle from 12 o'clock (in degree) - minutes granularity
	deg := (360 / 60) * (time.Minute())

	// (in radian)
	rad := (float64(2) * math.Pi) * (float64(deg) / 360.0)

	// calculate offsets from center (note: 90 is the second hand length)
	xOffset := math.Sin(rad) * 90
	yOffset := math.Cos(rad) * 90

	secondHandPoint := Point{
		X: (150.0 + xOffset),
		Y: (150.0 - yOffset),
	}
	return secondHandPoint

	// return Point{150, 150 - 90}
}

func secondsInRadians(time time.Time) float64 {
	return math.Pi
}
