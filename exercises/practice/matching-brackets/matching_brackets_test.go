package matchingbrackets

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
	for range b.N {
		for _, tt := range testCases {
			Bracket(tt.input)
		}
	}
}
