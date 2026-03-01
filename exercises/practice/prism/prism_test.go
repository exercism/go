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
	for range b.N {
		for _, tc := range testCases {
			FindSequence(tc.start, tc.prisms)
		}
	}
}
