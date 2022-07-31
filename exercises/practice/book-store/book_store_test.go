package bookstore

import (
	"testing"
)

func TestCost(t *testing.T) {
	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			actual := Cost(testCase.basket)
			if testCase.expected != actual {
				t.Errorf("Cost(%v) expected %d, got %d", testCase.basket, testCase.expected, actual)
			}
		})
	}
}

func BenchmarkCost(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			Cost(testCase.basket)
		}
	}
}
