package yacht

import (
	"testing"
)

var categoryNames = []string{"ones", "twos", "threes", "fours",
	"fives", "sixes", "Full House", "Four Of A Kind",
	"Little Straight", "Big Straight", "Choice", "Yacht"}

func TestCost(t *testing.T) {
	for _, testCase := range testCases {
		score := Score(testCase.dice, testCase.category)
		if testCase.expected != score {
			t.Fatalf("FAIL: %s\nScore %v as %s. expected %d, got %d",
				testCase.description, testCase.dice, categoryNames[testCase.category],
				testCase.expected, score)
		}
		t.Logf("PASS: %s", testCase.description)
	}
}

func BenchmarkCost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			Score(testCase.dice, testCase.category)
		}
	}
}
