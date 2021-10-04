package main

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const (
	secondHandLength = 90
	clockCenterX     = 150
	clockCenterY     = 150
)

func SecondHand(time time.Time) Point {
	secondHandPoint := secondHandPoint(time)

	// Scale to secondhand size (note: 90 is the second hand length)
	x := secondHandPoint.X * secondHandLength
	y := secondHandPoint.Y * secondHandLength

	// Flip (X-axis is pos on right side but Y-axis is neg on up side --- based on unit circle)
	x = x
	y = -y

	// Offset / translate
	x = x + clockCenterX
	y = y + clockCenterY

	return Point{x, y}
}

func secondsInRadians(time time.Time) float64 {
	return math.Pi / (30 / (float64(time.Second())))
}

func secondHandPoint(t time.Time) Point {
	rad := secondsInRadians(t)

	xPos := math.Sin(rad)
	yPos := math.Cos(rad)

	return Point{xPos, yPos}
}
