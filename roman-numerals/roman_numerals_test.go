package romannumerals

import "testing"

const testVersion = 1

// Retired testVersions
// (none) 313e3266c5fc18aca31a314b390bbc645956dbff

func TestRomanNumerals(t *testing.T) {
	if TestVersion != testVersion {
		t.Fatalf("Found TestVersion = %v, want %v", TestVersion, testVersion)
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
			t.Errorf("ToRomanNumeral(%d) should not return an error.", test.arabic)
			continue
		}
		if actual != test.roman {
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
