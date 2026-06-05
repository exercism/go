package etl

import (
	"maps"
	"testing"
)

func TestTransform(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Transform(tc.input); !maps.Equal(actual, tc.expected) {
				t.Fatalf("Transform(%v)\ngot: %v\nwant: %v", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkTransform(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			Transform(tc.input)
		}
	}
}
