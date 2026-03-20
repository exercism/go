package spiralmatrix

import (
	"reflect"
	"testing"
)

func TestSpiralMatrix(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := SpiralMatrix(tc.input)
			if !reflect.DeepEqual(got, tc.expected) {
				t.Fatalf("SpiralMatrix(%d)\n got: %v\nwant: %v", tc.input, got, tc.expected)
			}
		})
	}
}

func BenchmarkSpiralMatrix(b *testing.B) {
	for range b.N {
		for _, testCase := range testCases {
			SpiralMatrix(testCase.input)
		}
	}
}
