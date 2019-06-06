package darts

import (
	"testing"
)

func TestScore(t *testing.T) {
	for _, tc := range testCases {
		actual := Score(tc.x, tc.y)
		if actual != tc.expected {
			t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Score(tc.x, tc.y)
		}
	}
}
