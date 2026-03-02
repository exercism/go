package resistorcolorduo

import "testing"

func TestValue(t *testing.T) {
	for _, tc := range valueTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Value(tc.input)

			if actual != tc.expected {
				t.Fatalf("Value(%+q) = %d, want %d", tc.input, actual, tc.expected)
			}
		})
	}
}

// valueBench is intended to be used in BenchmarkValue to avoid compiler optimizations.
var valueBench int

func BenchmarkValue(b *testing.B) {
	for range b.N {
		for _, tc := range valueTestCases {
			valueBench = Value(tc.input)
		}
	}
}
