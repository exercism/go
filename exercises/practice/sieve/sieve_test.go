package sieve

import (
	"reflect"
	"testing"
)

func TestSieve(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Sieve(tc.limit)
			// use len() to allow either nil or empty list, because they are not equal by DeepEqual
			if len(actual) == 0 && len(tc.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("Sieve(%d)\n got:%#v\nwant:%#v", tc.limit, actual, tc.expected)
			}
		})
	}
}

func BenchmarkSieve(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Sieve(tc.limit)
		}
	}
}
