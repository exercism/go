package prime

import (
	"reflect"
	"sort"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	for _, test := range tests {
		actual := Factors(test.input)
		sort.Slice(actual, ascending(actual))
		sort.Slice(test.expected, ascending(test.expected))
		if !reflect.DeepEqual(actual, test.expected) {
			t.Fatalf("FAIL %s\nFactors(%d) = %#v;\nexpected %#v",
				test.description, test.input,
				actual, test.expected)
		}
		t.Logf("PASS %s", test.description)
	}
}

func BenchmarkPrimeFactors(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			Factors(test.input)
		}
	}
}

func ascending(list []int64) func(int, int) bool {
	return func(i, j int) bool {
		return list[i] < list[j]
	}
}
