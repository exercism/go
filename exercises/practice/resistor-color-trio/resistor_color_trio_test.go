package resistorcolortrio

import "testing"

func TestLabel(t *testing.T) {
	for _, tc := range labelTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Label(tc.input)

			if actual != tc.expected {
				t.Fatalf("Label(%+q) = %q, want %q", tc.input, actual, tc.expected)
			}
		})
	}
}

// labelBench is intended to be used in BenchmarkLabel to avoid compiler optimizations.
var labelBench string

func BenchmarkLabel(b *testing.B) {
	for range b.N {
		for _, tc := range labelTestCases {
			labelBench = Label(tc.input)
		}
	}
}
