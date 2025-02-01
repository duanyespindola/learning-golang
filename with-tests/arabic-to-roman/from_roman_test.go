package main

import (
	"fmt"
	"testing"
)

func TestRomanToArabic(t *testing.T) {
	tests := []struct {
		arabic int
		roman  string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{14, "XIV"},
		{18, "XVIII"},
		{20, "XX"},
		{39, "XXXIX"},
		{40, "XL"},
		{50, "L"},
		{90, "XC"},
		{100, "C"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Number %d", test.arabic), func(t *testing.T) {
			got := RomanToArabic(test.roman)
			if got != test.arabic {
				t.Errorf("got %d, want %d", got, test.arabic)
			}
		})
	}
}

var allRomanNumerals = map[string]int{
	"M": 1000,
	"D": 500,
	"C": 100,
	"L": 50,
	"X": 10,
	"V": 5,
	"I": 1,
}

func RomanToArabic(roman string) int {
	if roman == "" {
		return 0
	}
	if len(roman) == 1 {
		return allRomanNumerals[roman]
	}

	first := string(roman[0])
	second := string(roman[1])
	if allRomanNumerals[first] < allRomanNumerals[second] {
		return RomanToArabic(roman[1:]) - allRomanNumerals[first]
	}

	return RomanToArabic(roman[1:]) + allRomanNumerals[first]
}
