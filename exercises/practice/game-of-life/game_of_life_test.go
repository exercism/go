package gameoflife

import (
	"slices"
	"testing"
)

func TestTick(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Tick(tc.input); !slices.EqualFunc(actual, tc.expected, slices.Equal) {
				t.Fatalf("Tick(%v) = %v, want: %v", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkTick(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			Tick(tc.input)
		}
	}
}
