package romannumerals

import (
	"testing"
)

type romanNumeralTest struct {
	arabic   int
	roman    string
	hasError bool
}

var romanNumeralTests = []romanNumeralTest{
	{1, "I", false},
	{2, "II", false},
	{3, "III", false},
	{4, "IV", false},
	{5, "V", false},
	{6, "VI", false},
	{9, "IX", false},
	{27, "XXVII", false},
	{48, "XLVIII", false},
	{59, "LIX", false},
	{93, "XCIII", false},
	{141, "CXLI", false},
	{163, "CLXIII", false},
	{402, "CDII", false},
	{575, "DLXXV", false},
	{911, "CMXI", false},
	{1024, "MXXIV", false},
	{3000, "MMM", false},
	{3001, "", false}, // somewhat arbitrary cut-off. Perhaps 10,000 would be better?
	{0, "", true},
	{-1, "", true},
}

func TestRomanNumerals(t *testing.T) {
	for _, test := range romanNumeralTests {
		actual, err := ToRomanNumeral(test.arabic)
		if err == nil && test.hasError {
			t.Errorf("ToRomanNumberal(%d) should return an error.", test.arabic)
		}
		if actual != test.roman {
			t.Errorf("ToRomanNumeral(%d): expected %s, actual %s", test.arabic, test.roman, actual)
		}
	}
}

func BenchmarkRomanNumerals(b *testing.B) {
	b.StopTimer()
	for _, test := range romanNumeralTests {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			ToRomanNumeral(test.arabic)
		}

		b.StopTimer()
	}
}
