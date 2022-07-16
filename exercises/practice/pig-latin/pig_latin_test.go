package piglatin

import (
	"testing"
)

func TestPigLatin(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Sentence(tc.input); actual != tc.expected {
				t.Fatalf("Sentence(%q) = %q, want %q", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkSentence(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Sentence(tc.input)
		}
	}
}
