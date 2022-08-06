package romannumerals

import "testing"

func TestRomanNumerals(t *testing.T) {
	for _, tc := range append(validRomanNumeralTests, invalidRomanNumeralTests...) {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := ToRomanNumeral(tc.input)
			switch {
			case tc.expectError:
				if err == nil {
					t.Fatalf("ToRomanNumeral(%d) expected error, got: %s", tc.input, actual)
				}
			case err != nil:
				t.Fatalf("ToRomanNumeral(%d) returned error: %v, want: %q", tc.input, err, tc.expected)
			case actual != tc.expected:
				t.Fatalf("ToRomanNumeral(%d) = %q, want: %q", tc.input, actual, tc.expected)
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

var invalidRomanNumeralTests = []romanNumeralTest{
	{
		description: "0 is out of range",
		input:       0,
		expectError: true,
	},
	{
		description: "-1 is out of range",
		input:       -1,
		expectError: true,
	},
	{
		description: "3001 is out of range",
		input:       3001,
		expectError: true,
	},
}
