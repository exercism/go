package grains

import (
	"testing"
)

func TestSquare(t *testing.T) {
	for _, tc := range squareTests {
		t.Run(tc.description, func(t *testing.T) {
			actualVal, actualErr := Square(tc.input)
			if tc.expectError {
				if actualErr == nil {
					t.Errorf("Square(%d) expected an error, got: %d", tc.input, actualVal)
				}
			} else {
				if actualErr != nil {
					t.Errorf("Square(%d) expected %d, but got error: %v", tc.input, actualVal, actualErr)
				} else if actualVal != tc.expectedVal {
					t.Errorf("Square(%d) = %d, want %d", tc.input, actualVal, tc.expectedVal)
				}
			}
		})
	}
}

func TestTotal(t *testing.T) {
	var expected uint64 = 18446744073709551615
	if actual := Total(); actual != expected {
		t.Errorf("Total() = %d, want:%d", actual, expected)
	}
}

func BenchmarkSquare(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	for i := 0; i < b.N; i++ {
		for _, test := range squareTests {
			Square(test.input)
		}
	}
}

func BenchmarkTotal(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		Total()
	}
}
