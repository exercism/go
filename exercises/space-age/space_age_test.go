package space

import (
	"math"
	"testing"
)

func TestAge(t *testing.T) {
	const precision = 0.01
	for _, tc := range testCases {
		if actual := Age(tc.seconds, tc.planet); math.Abs(actual-tc.expected) > precision {
			t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkAge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Age(tc.seconds, tc.planet)
		}
	}
}
