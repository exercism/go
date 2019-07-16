package darts

import (
	"testing"
)

func TestScore(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Score(tc.x, tc.y)
			if actual != tc.expected {
				t.Errorf("FAIL: %s\nExpected: %#v\nActual: %#v", tc.description, tc.expected, actual)
			}
		})
	}
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Score(tc.x, tc.y)
		}
	}
}
