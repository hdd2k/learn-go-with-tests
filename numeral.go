package main

import (
	"strings"
)

type RomanNum struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNum

var RomanNums = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func (r RomanNumerals) ValueOf(symbols ...byte) int {
	sym := string(symbols)
	for _, s := range r {
		if sym == s.Symbol {
			return s.Value
		}
	}
	return 0
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, numeral := range RomanNums {
		for arabic >= numeral.Value {
			result.WriteString(numeral.Symbol)
			arabic -= numeral.Value
		}
	}
	return result.String()
}

func ConvertToArabic(roman string) int {
	total := 0
	for i := 0; i < len(roman); i++ {
		currSym := roman[i]

		// lookahead to next char if possible
		if couldBeSubtracted(i, len(roman), currSym) {
			nextSym := roman[i+1]
			// convert combined string by looking up
			val := RomanNums.ValueOf(currSym, nextSym)
			if val != 0 {
				total += val
				i++
			} else {
				total += RomanNums.ValueOf(currSym)
			}
		} else {
			total += (RomanNums.ValueOf(currSym))
		}
	}

	return total

}

func couldBeSubtracted(index int, allRomanLen int, currSym uint8) bool {
	return (index+1 < allRomanLen) && ((currSym == 'I') || (currSym == 'X') || (currSym == 'C'))

}
