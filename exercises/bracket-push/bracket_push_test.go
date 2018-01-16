package brackets

import (
	"testing"
)

func TestBracket(t *testing.T) {
	for _, tt := range testCases {
		actual, err := Bracket(tt.input)
		// We don't expect errors for any of the test cases
		if err != nil {
			var _ error = err
			t.Fatalf("Bracket(%q) returned error %q.  Error not expected.", tt.input, err)
		}
		if actual != tt.expected {
			t.Fatalf("Bracket(%q) was expected to return %v but returned %v.",
				tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkBracket(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			Bracket(tt.input)
		}
	}
}
