package differenceofsquares

import "testing"

func TestDifferenceOfSquares(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if got := tc.fn(tc.input); got != tc.expected {
				t.Fatalf("%s(%d) = %d, want: %d", tc.fnName, tc.input, got, tc.expected)
			}
		})
	}
}

func BenchmarkSquareOfSum(b *testing.B) {
	for range b.N {
		SquareOfSum(100)
	}
}

func BenchmarkSumOfSquares(b *testing.B) {
	for range b.N {
		SumOfSquares(100)
	}
}

func BenchmarkDifference(b *testing.B) {
	for range b.N {
		Difference(100)
	}
}
