package diffsquares

import (
	"fmt"
	"testing"
)

var tests = []struct{ input, squareOfSum, sumOfSquares int }{
	{input: 5, squareOfSum: 225, sumOfSquares: 55},
	{input: 10, squareOfSum: 3025, sumOfSquares: 385},
	{input: 100, squareOfSum: 25502500, sumOfSquares: 338350},
}

func TestSquareOfSum(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("Square of sum from 1 to %d", test.input), func(t *testing.T) {
			if got := SquareOfSum(test.input); got != test.squareOfSum {
				t.Fatalf("SquareOfSum(%d) = %d, want: %d", test.input, got, test.squareOfSum)
			}
		})
	}
}

func TestSumOfSquares(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("Sum of squares from 1 to %d", test.input), func(t *testing.T) {
			if got := SumOfSquares(test.input); got != test.sumOfSquares {
				t.Fatalf("SumOfSquares(%d) = %d, want: %d", test.input, got, test.sumOfSquares)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	for _, test := range tests {
		t.Run(fmt.Sprintf("Difference of SquareOfSum and SumOfSquares of value %d", test.input), func(t *testing.T) {
			want := test.squareOfSum - test.sumOfSquares
			if got := Difference(test.input); got != want {
				t.Fatalf("Difference(%d) = %d, want: %d", test.input, got, want)
			}
		})
	}
}

// Benchmark functions on just a single number (100, from the original PE problem)
// to avoid overhead of iterating over tests.
func BenchmarkSquareOfSum(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		SquareOfSum(100)
	}
}

func BenchmarkSumOfSquares(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		SumOfSquares(100)
	}
}

func BenchmarkDifference(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		Difference(100)
	}
}
