package romannumerals

import "testing"

const targetTestVersion = 3

func TestRomanNumerals(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
	tc := append(romanNumeralTests, []romanNumeralTest{
		{0, "", true},
		{-1, "", true},
		{4000, "", true},
		{3999, "MMMCMXCIX", false},
	}...)
	for _, test := range tc {
		actual, err := ToRomanNumeral(test.arabic)
		if err == nil && test.hasError {
			t.Errorf("ToRomanNumeral(%d) should return an error.", test.arabic)
			continue
		}
		if err != nil && !test.hasError {
			var _ error = err
			t.Errorf("ToRomanNumeral(%d) should not return an error.", test.arabic)
			continue
		}
		if !test.hasError && actual != test.roman {
			t.Errorf("ToRomanNumeral(%d): %s, expected %s", test.arabic, actual, test.roman)
		}
	}
}

func BenchmarkRomanNumerals(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range romanNumeralTests {
			ToRomanNumeral(test.arabic)
		}
	}
}
