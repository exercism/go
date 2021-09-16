package brackets

import (
	"testing"
)

func TestBracket(t *testing.T) {
	for _, tt := range testCases {
		actual := Bracket(tt.input)
		if actual != tt.expected {
			t.Fatalf("Bracket(%q) was expected to return %v but returned %v.",
				tt.input, tt.expected, actual)
		}
	}
}

func BenchmarkBracket(b *testing.B) {
	if testing.Short() {
		t.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			Bracket(tt.input)
		}
	}
}
