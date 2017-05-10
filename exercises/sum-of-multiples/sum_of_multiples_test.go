package summultiples

import "testing"

const targetTestVersion = 1

var varTests = []struct {
	divisors []int
	limit    int
	sum      int
}{
	// converted from x-common by hand
	// missing: 4/6 to 15, 5/6/7 to 150, 5/25 to 51, 1 to 100
	// extra: 1/1 (duplicate factor?)
	// x-common version: 0.5.0
	{[]int{3, 5}, 1, 0},
	{[]int{3, 5}, 4, 3},
	{[]int{3, 5}, 10, 23},
	{[]int{3, 5}, 100, 2318},
	{[]int{3, 5}, 1000, 233168},
	{[]int{7, 13, 17}, 20, 51},
	{[]int{43, 47}, 10000, 2203160},
	{[]int{5, 10, 12}, 10000, 13331672},
	{[]int{1, 1}, 10000, 49995000},
	{[]int{}, 10000, 0},
}

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

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
	for i := 0; i < b.N; i++ {
		for _, test := range varTests {
			SumMultiples(test.limit, test.divisors...)
		}
	}
}
