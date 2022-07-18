package yacht

import (
	"testing"
)

func TestScore(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Score(tc.dice, tc.category)
			if tc.expected != actual {
				t.Fatalf("Score(%#v,%q) = %d, want:%d", tc.dice, tc.category, actual, tc.expected)
			}
		})
	}
}

func BenchmarkScore(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			Score(testCase.dice, testCase.category)
		}
	}
}
