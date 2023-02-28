package resistorcolorduo

import "testing"

func TestValue(t *testing.T) {
	for _, tc := range valueTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Value(tc.input)

			if actual != tc.expected {
				t.Fatalf("Value(%+q): expected %d, actual %d", tc.input, tc.expected, actual)
			}
		})
	}
}

// valueBench is intended to be used in BenchmarkValue to avoid compiler optimizations.
var valueBench int

func BenchmarkValue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range valueTestCases {
			valueBench = Value(tc.input)
		}
	}
}
