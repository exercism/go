package prism

import (
	"slices"
	"testing"
)

func TestFindSequence(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := FindSequence(tc.start, tc.prisms); !slices.Equal(actual, tc.expected) {
				t.Fatalf("FindSequence(%v), %v = %v, want: %v", tc.start, tc.prisms, actual, tc.expected)
			}
		})
	}
}

func BenchmarkFindSequence(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			FindSequence(tc.start, tc.prisms)
		}
	}
}
