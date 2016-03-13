package summultiples

import "testing"

var test35 = []struct {
	limit int
	sum   int
}{
	{1, 0},
	{4, 3},
	{10, 23},
	{100, 2318},
	{1000, 233168},
}

var varTests = []struct {
	divisors []int
	limit    int
	sum      int
}{
	{[]int{7, 13, 17}, 20, 51},
	{[]int{43, 47}, 10000, 2203160},
	{[]int{5, 10, 12}, 10000, 13331672},
	{[]int{1, 1}, 10000, 49995000},
	// Note: The following test case deviates from the README.
	// The README specifies some rather odd defaults, whereas
	// this has the more logical approach of not implementing any
	// defaults, which causes the resulting sum to be zero.
	// See discussion in:
	// https://github.com/exercism/xgo/issues/256 and
	// https://github.com/exercism/x-common/issues/198
	{[]int{}, 10000, 0},
}

func Test35(t *testing.T) {
	sum35 := MultipleSummer(3, 5)
	for _, test := range test35 {
		s := sum35(test.limit)
		if s != test.sum {
			t.Fatalf("Sum to %d returned %d, want %d.", test.limit, s, test.sum)
		}
	}
}

func TestVar(t *testing.T) {
	for _, test := range varTests {
		sv := MultipleSummer(test.divisors...)
		s := sv(test.limit)
		if s != test.sum {
			t.Fatalf("Sum of multiples of %v to %d returned %d, want %d.",
				test.divisors, test.limit, s, test.sum)
		}
	}
}

func Benchmark35(b *testing.B) {
	sum35 := MultipleSummer(3, 5)
	b.ResetTimer() // bench just the sum function
	for i := 0; i < b.N; i++ {
		for _, test := range test35 {
			sum35(test.limit)
		}
	}
}

func BenchmarkVar(b *testing.B) {
	// bench combined time to bind sum function and call it.
	for i := 0; i < b.N; i++ {
		for _, test := range varTests {
			MultipleSummer(test.divisors...)(test.limit)
		}
	}
}
