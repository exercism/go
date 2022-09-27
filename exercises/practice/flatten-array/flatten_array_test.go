package flatten

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Flatten(tc.input)

			// Clarify failures where an empty slice is expected, but a nil slice is given.
			if actual == nil && tc.expected != nil && len(tc.expected) == 0 {
				t.Fatalf("Flatten(%v) = %v (nil slice), want: %v (empty slice)", tc.input, actual, tc.expected)
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("Flatten(%v) = %v, want: %v", tc.input, &actual, tc.expected)
			}
		})
	}
}

func BenchmarkFlatten(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Flatten(tc.input)
		}
	}
}
