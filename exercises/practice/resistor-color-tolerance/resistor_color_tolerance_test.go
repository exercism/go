package resistorcolortolerance

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

func BenchmarkLabel(b *testing.B) {
	for b.Loop() {
		for _, tc := range labelTestCases {
			Label(tc.input)
		}
	}
}
