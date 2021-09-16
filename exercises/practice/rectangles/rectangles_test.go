package rectangles

import (
	"testing"
)

func TestRectangles(t *testing.T) {
	for _, tc := range testCases {
		if actual := Count(tc.input); actual != tc.expected {
			t.Fatalf("FAIL: %s\nExpected: %#v\nActual: %#v", tc.description, tc.expected, actual)
		}
		t.Logf("PASS: %s", tc.description)
	}
}

func BenchmarkRectangles(b *testing.B) {
	if testing.Short() {
		t.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Count(tc.input)
		}
	}
}
