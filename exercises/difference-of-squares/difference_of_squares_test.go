package diffsquares

import "testing"

const targetTestVersion = 1

var tests = []struct{ n, sqOfSums, sumOfSq int }{
	// converted from x-common by hand
	// Only needs a test of 0 to be compliant with 1.0.0.
	// x-common version: 0.9.0
	{5, 225, 55},
	{10, 3025, 385},
	{100, 25502500, 338350},
}

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func TestSquareOfSums(t *testing.T) {
	for _, test := range tests {
		if s := SquareOfSums(test.n); s != test.sqOfSums {
			t.Fatalf("SquareOfSums(%d) = %d, want %d", test.n, s, test.sqOfSums)
		}
	}
}

func TestSumOfSquares(t *testing.T) {
	for _, test := range tests {
		if s := SumOfSquares(test.n); s != test.sumOfSq {
			t.Fatalf("SumOfSquares(%d) = %d, want %d", test.n, s, test.sumOfSq)
		}
	}
}

func TestDifference(t *testing.T) {
	for _, test := range tests {
		want := test.sqOfSums - test.sumOfSq
		if s := Difference(test.n); s != want {
			t.Fatalf("Difference(%d) = %d, want %d", test.n, s, want)
		}
	}
}

// Benchmark functions on just a single number (100, from the original PE problem)
// to avoid overhead of iterating over tests.
func BenchmarkSquareOfSums(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SquareOfSums(100)
	}
}

func BenchmarkSumOfSquares(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumOfSquares(100)
	}
}

func BenchmarkDifference(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Difference(100)
	}
}
