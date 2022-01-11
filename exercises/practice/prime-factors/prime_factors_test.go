package prime

import (
	"sort"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	for _, test := range tests {
		actual := Factors(test.input)
		sort.Slice(actual, ascending(actual))
		sort.Slice(test.expected, ascending(test.expected))
		if !slicesEqual(actual, test.expected) {
			t.Fatalf("FAIL %s\nFactors(%d) = %#v;\nexpected %#v",
				test.description, test.input,
				actual, test.expected)
		}
		t.Logf("PASS %s", test.description)
	}
}

func BenchmarkPrimeFactors(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Factors(test.input)
		}
	}
}

func slicesEqual(a, b []int64) bool {
	if len(a) != len(b) {
		return false
	}

	if len(a) == 0 {
		return true
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
