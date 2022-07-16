package brackets

import (
	"testing"
)

func TestBracket(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Bracket(tc.input)
			if actual != tc.expected {
				t.Fatalf("Bracket(%q) = %t, want: %t", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkBracket(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			Bracket(tt.input)
		}
	}
}
