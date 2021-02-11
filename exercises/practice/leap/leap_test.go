package leap

import "testing"

func TestLeapYears(t *testing.T) {
	for _, test := range testCases {
		observed := IsLeapYear(test.year)
		if observed != test.expected {
			t.Fatalf("IsLeapYear(%d) = %t, want %t (%s)",
				test.year, observed, test.expected, test.description)
		}
	}
}

// Benchmark 400 year interval to get fair weighting of different years.
func Benchmark400(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for y := 1600; y < 2000; y++ {
			IsLeapYear(y)
		}
	}
}
