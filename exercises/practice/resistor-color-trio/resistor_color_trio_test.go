package resistorcolortrio

import "testing"

func TestLabel(t *testing.T) {
	for _, tc := range labelTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Label(tc.input)

			if actual != tc.expected {
				t.Fatalf("Label(%+q): expected %q, actual %q", tc.input, tc.expected, actual)
			}
		})
	}
}

// labelBench is intended to be used in BenchmarkLabel to avoid compiler optimizations.
var labelBench string

func BenchmarkLabel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range labelTestCases {
			labelBench = Label(tc.input)
		}
	}
}
