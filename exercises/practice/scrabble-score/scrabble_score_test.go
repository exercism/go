package scrabblescore

import "testing"

func TestScore(t *testing.T) {
	for _, tc := range scrabbleScoreTests {
		t.Run(tc.description, func(t *testing.T) {
			if actual := Score(tc.input); actual != tc.expected {
				t.Errorf("Score(%q) = %d, want:%d", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkScore(b *testing.B) {
	for range b.N {
		for _, test := range scrabbleScoreTests {
			Score(test.input)
		}
	}
}
