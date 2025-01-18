package prime

import (
	"slices"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Factors(tc.input)
			slices.Sort(actual)
			slices.Sort(tc.expected)
			if !slices.Equal(actual, tc.expected) {
				t.Fatalf("Factors(%d)\n got:%#v\nwant:%#v", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkPrimeFactors(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			Factors(test.input)
		}
	}
}
