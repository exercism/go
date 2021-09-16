package bookstore

import (
	"testing"
)

func TestCost(t *testing.T) {
	for _, testCase := range testCases {
		cost := Cost(testCase.basket)
		if testCase.expected != cost {
			t.Fatalf("FAIL: %s\n\tCost(%v) expected %v, got %v",
				testCase.description, testCase.basket, testCase.expected, cost)
		}
		t.Logf("PASS: %s", testCase.description)
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
