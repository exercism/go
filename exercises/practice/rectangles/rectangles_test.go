package rectangles

import (
	"testing"
)

func TestRectangles(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Count(tc.input); actual != tc.expected {
				t.Fatalf("Count(%#v) = %d, want: %d", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkRectangles(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			Count(tc.input)
		}
	}
}
