package yacht

import (
	"testing"
)

func TestScore(t *testing.T) {
	for _, testCase := range testCases {
		score := Score(testCase.dice, testCase.category)
		if testCase.expected != score {
			t.Fatalf("FAIL: %s\nScore %v as %s. expected %d, got %d",
				testCase.description, testCase.dice, testCase.category,
				testCase.expected, score)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

func BenchmarkScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			Score(testCase.dice, testCase.category)
		}
	}
}
