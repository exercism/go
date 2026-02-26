package squareroot

import "testing"

func TestSquareRoot(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual, err := SquareRoot(tc.input); err != nil {
				t.Fatalf("SquareRoot(%d) returned error: %v, want: %q", tc.input, err, tc.expected)
			} else if actual != tc.expected {
				t.Fatalf("SquareRoot(%d) = %d, want: %d", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkSquareRoot(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			SquareRoot(tc.input)
		}
	}
}
