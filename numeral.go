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

func (r RomanNumerals) ValueOf(otherStr string) int {
	for _, s := range r {
		if otherStr == s.Symbol {
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
	// var result int
	// for range roman {
	// 	result += 1
	// }
	// return result

	total := 0
	for i := 0; i < len(roman); i++ {
		currSym := roman[i]

		// lookahead to next char if possible
		if (i+1 < len(roman)) && (currSym == 'I') {
			nextSym := roman[i+1]
			// create combined string
			potentialNum := string([]byte{currSym, nextSym})
			// convert combined string by looking up
			val := RomanNums.ValueOf(potentialNum)
			if val != 0 {
				total += val
				i++
			} else {
				total++
			}
		} else {
			total++
		}
	}

	return total

}
