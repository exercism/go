package space

import (
	"math"
	"testing"
)

func TestAge(t *testing.T) {
	const precision = 0.01
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Age(tc.seconds, tc.planet)
			if math.Abs(actual-tc.expected) > precision {
				t.Fatalf("Age(%f, %v) = %f, want: %f", tc.seconds, tc.planet, actual, tc.expected)
			}
		})
	}
}

func BenchmarkAge(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Age(tc.seconds, tc.planet)
		}
	}
}
