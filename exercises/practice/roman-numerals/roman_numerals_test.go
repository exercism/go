package romannumerals

import "testing"

func TestRomanNumerals(t *testing.T) {
	for _, tc := range validRomanNumeralTests {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := ToRomanNumeral(tc.input)
			if err != nil {
				// expect no error for all valid tests cases (canonical-data.json contains only valid cases)
				t.Fatalf("ToRomanNumeral(%d) returned error: %v, want: %q", tc.input, err, tc.expected)
			}
			if actual != tc.expected {
				t.Fatalf("ToRomanNumeral(%d) = %q, want: %q", tc.input, actual, tc.expected)
			}
		})

	}
}

func TestRomanNumeralsInvalid(t *testing.T) {
	invalidRomanNumeralTests := []romanNumeralTest{
		{description: "0 is out of range", input: 0},
		{description: "-1 is out of range", input: -1},
		{description: "4000 is out of range", input: 4000},
	}
	for _, tc := range invalidRomanNumeralTests {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := ToRomanNumeral(tc.input)
			if err == nil {
				t.Fatalf("ToRomanNumeral(%d) expected error, got: %q", tc.input, actual)
			}
		})
	}
}

func BenchmarkRomanNumerals(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range validRomanNumeralTests {
			ToRomanNumeral(tc.input)
		}
	}
}
