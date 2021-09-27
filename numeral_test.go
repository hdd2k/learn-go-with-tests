package main

import (
	"fmt"
	"testing"
)

var cases = []struct {
	Desc   string
	Arabic int
	Roman  string
}{
	{"1 converts to I", 1, "I"},
	{"2 converts to II", 2, "II"},
	{"3 converts to III", 3, "III"},
	{"4 converts to IV", 4, "IV"},
	{"5 converts to V", 5, "V"},
	{"6 converts to VI", 6, "VI"},
	{"7 converts to VII", 7, "VII"},
	{"8 converts to VIII", 8, "VIII"},
	{"9 converts to IX", 9, "IX"},
	{"10 converts to X", 10, "X"},
	{"14 gets converted to XIV", 14, "XIV"},
	{"18 gets converted to XVIII", 18, "XVIII"},
	{"20 gets converted to XX", 20, "XX"},
	{"39 gets converted to XXXIX", 39, "XXXIX"},
	{"40 gets converted to XL", 40, "XL"},
	{"47 gets converted to XLVII", 47, "XLVII"},
	{"49 gets converted to XLIX", 49, "XLIX"},
	{"50 gets converted to L", 50, "L"},
	{"100 gets converted to L", 100, "C"},
	{"90 gets converted to L", 90, "XC"},
	{"400 gets converted to L", 400, "CD"},
	{"500 gets converted to L", 500, "D"},
	{"900 gets converted to L", 900, "CM"},
	{"1000 gets converted to L", 1000, "M"},
	{"1984 gets converted to L", 1984, "MCMLXXXIV"},
	{"3999 gets converted to L", 3999, "MMMCMXCIX"},
	{"2014 gets converted to L", 2014, "MMXIV"},
	{"1006 gets converted to L", 1006, "MVI"},
	{"798 gets converted to L", 798, "DCCXCVIII"},
}

func TestNumeral(t *testing.T) {
	for _, c := range cases {
		t.Run(c.Desc, func(t *testing.T) {
			want := c.Roman
			got := ConvertToRoman(c.Arabic)
			if want != got {
				t.Errorf("got %q want %q", got, want)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, c := range cases[:4] {
		t.Run(fmt.Sprintf("%s gets converted to %d", c.Roman, c.Arabic), func(t *testing.T) {
			want := c.Arabic
			got := ConvertToArabic(c.Roman)
			if want != got {
				t.Errorf("got %d want %d", got, want)
			}
		})
	}
}
