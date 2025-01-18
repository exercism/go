package darts

import (
	"testing"
)

func TestScore(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Score(tc.x, tc.y)
			if actual != tc.expected {
				t.Errorf("Score(%#v, %#v) = %#v, want: %#v", tc.x, tc.y, actual, tc.expected)
			}
		})
	}
}

func BenchmarkScore(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Score(tc.x, tc.y)
		}
	}
}
