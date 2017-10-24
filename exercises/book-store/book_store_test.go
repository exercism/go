package bookstore

import (
	"testing"
)

var testCases = []struct {
	description string
	basket      []int
	expected    float64
}{
	{
		description: "only a single book",
		basket:      []int{1},
		expected:    8.00,
	},
	{
		description: "two of the same book",
		basket:      []int{2, 2},
		expected:    16.00,
	},
	{
		description: "empty basket",
		basket:      []int{},
		expected:    0.00,
	},
	{
		description: "two different books",
		basket:      []int{1, 2},
		expected:    15.20,
	},
	{
		description: "four different books",
		basket:      []int{1, 2, 3, 4},
		expected:    25.60,
	},
	{
		description: "five different books",
		basket:      []int{1, 2, 3, 4, 5},
		expected:    30.00,
	},
	{
		description: "two groups of four is cheaper than group of five plus group of three",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 5},
		expected:    51.20,
	},
	{
		description: "groups of four plus group of two is is cheaper than two groups of three",
		basket:      []int{1, 1, 2, 2, 3, 4},
		expected:    40.80,
	},
	{
		description: "two each of first 4 books and 1 copy of each of rest",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5},
		expected:    55.60,
	},
	{
		description: "two copies of each book",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5},
		expected:    60.00,
	},
	{
		description: "three copies of first book and 2 each fo remaining",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 1},
		expected:    68.00,
	},
	{
		description: "three each of first 2 books and 2 each of remaining books",
		basket:      []int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 1, 2},
		expected:    75.20,
	},
}

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
	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			Cost(testCase.basket)
		}
	}
}
