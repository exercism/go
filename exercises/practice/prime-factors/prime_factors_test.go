package prime

import (
	"sort"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Factors(tc.input)
			sort.Slice(actual, ascending(actual))
			sort.Slice(tc.expected, ascending(tc.expected))
			if !slicesEqual(actual, tc.expected) {
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

func slicesEqual(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func ascending(list []int64) func(int, int) bool {
	return func(i, j int) bool {
		return list[i] < list[j]
	}
}
