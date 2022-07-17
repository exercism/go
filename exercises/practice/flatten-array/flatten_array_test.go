package flatten

import (
	"reflect"
	"testing"
)

func TestFlatten(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Flatten(tc.input); !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Flatten(%v) = %v, want: %v", tc.input, actual, tc.expected)
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
