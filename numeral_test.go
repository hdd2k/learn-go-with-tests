package main

import (
	"testing"
)

func TestNumeral(t *testing.T) {

	cases := []struct {
		Desc   string
		Arabic int
		Want   string
	}{
		{
			"1 converts to I",
			1,
			"I",
		},
		{
			"2 converts to II",
			2,
			"II",
		},
		{
			"3 converts to III",
			3,
			"III",
		},
		{
			"4 converts to IV",
			4,
			"IV",
		},
		{
			"5 converts to V",
			5,
			"V",
		},
		{
			"6 converts to VI",
			6,
			"VI",
		},
		{
			"7 converts to VII",
			7,
			"VII",
		},
		{
			"8 converts to VIII",
			8,
			"VIII",
		},
	}

	for _, c := range cases {
		t.Run(c.Desc, func(t *testing.T) {
			want := c.Want
			got := ConvertToRoman(c.Arabic)
			if want != got {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}
}
