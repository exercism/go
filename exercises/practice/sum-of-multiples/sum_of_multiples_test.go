package summultiples

import "testing"

func TestSumMultiples(t *testing.T) {
	for _, test := range varTests {
		s := SumMultiples(test.limit, test.divisors...)
		if s != test.sum {
			t.Fatalf("Sum of multiples of %v to %d returned %d, want %d.",
				test.divisors, test.limit, s, test.sum)
		}
	}
}

func BenchmarkSumMultiples(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range varTests {
			SumMultiples(test.limit, test.divisors...)
		}
	}
}
