package prime

// Return prime factors in increasing order

import (
	"reflect"
	"testing"
)

func TestPrimeFactors(t *testing.T) {
	for _, test := range tests {
		actual := Factors(test.input)
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
