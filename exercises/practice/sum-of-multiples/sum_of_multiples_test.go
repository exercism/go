package summultiples

import "testing"

func TestSumMultiples(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := SumMultiples(tc.limit, tc.divisors...)
			if actual != tc.expected {
				t.Fatalf("SumMultiples(%d, %#v) = %d, want: %d", tc.limit, tc.divisors, actual, tc.expected)
			}
		})
	}
}

func BenchmarkSumMultiples(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			SumMultiples(tc.limit, tc.divisors...)
		}
	}
}
