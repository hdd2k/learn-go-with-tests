package main

import (
	"bytes"
	"encoding/xml"
	"math"
	"testing"
	"time"
)

// every clock has a centre of (150, 150)
// the hour hand is 50 long
// the minute hand is 80 long
// the second hand is 90 long.

type Svg struct {
	XMLName xml.Name `xml:"svg"`
	Text    string   `xml:",chardata"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  struct {
		Text  string `xml:",chardata"`
		Cx    string `xml:"cx,attr"`
		Cy    string `xml:"cy,attr"`
		R     string `xml:"r,attr"`
		Style string `xml:"style,attr"`
	} `xml:"circle"`
	Line []struct {
		Text  string `xml:",chardata"`
		X1    string `xml:"x1,attr"`
		Y1    string `xml:"y1,attr"`
		X2    string `xml:"x2,attr"`
		Y2    string `xml:"y2,attr"`
		Style string `xml:"style,attr"`
	} `xml:"line"`
}

// func TestSecondHandAtMidnight(t *testing.T) {
// 	time := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

// 	want := Point{X: 150, Y: 150 - 90}
// 	got := SecondHand(time)

// 	if want != got {
// 		t.Errorf("Got %v want %v", got, want)
// 	}
// }

// func TestSecondHandAt30Seconds(t *testing.T) {
// 	time := time.Date(1337, time.January, 1, 0, 0, 30, 0, time.UTC)

// 	want := Point{X: 150, Y: 150 + 90}
// 	got := SecondHand(time)

// 	if want != got {
// 		t.Errorf("Got %v want %v", got, want)
// 	}
// }

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

func TestSVGWriterAtMidnight(t *testing.T) {
	tm := time.Date(1337, time.January, 1, 0, 0, 0, 0, time.UTC)

	b := bytes.Buffer{}
	SVGWriter(&b, tm)

	svg := Svg{}
	xml.Unmarshal(b.Bytes(), &svg)

	x2 := "150.000"
	y2 := "60.000"

	for _, line := range svg.Line {
		if line.X2 == x2 && line.Y2 == y2 {
			return
		}
	}

	t.Errorf("Expected to find second hand pattern with x2 of %+v and y2 of %+v, in SVG output %v", x2, y2, b.String())
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
