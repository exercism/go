package killersudokuhelper

import (
	"slices"
	"testing"
)

func TestCombinations(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Combinations(tc.sum, tc.size, tc.exclude); !slices.EqualFunc(actual, tc.expected, slices.Equal) {
				t.Fatalf("Combinations(%d, %d, %#v) = %#v, want: %#v", tc.sum, tc.size, tc.exclude, actual, tc.expected)
			}
		})
	}
}

func BenchmarkCombinations(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			Combinations(tc.sum, tc.size, tc.exclude)
		}
	}
}
