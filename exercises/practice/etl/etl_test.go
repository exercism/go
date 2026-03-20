package etl

import (
	"maps"
	"testing"
)

func TestTransform(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.description, func(t *testing.T) {
			if actual := Transform(tt.input); !maps.Equal(actual, tt.expected) {
				t.Fatalf("Transform(%v)\ngot: %v\nwant: %v", tt.input, actual, tt.expected)
			}
		})
	}
}

func BenchmarkTransform(b *testing.B) {
	for range b.N {
		for _, tt := range testCases {
			Transform(tt.input)
		}
	}
}
