package leap

import "testing"

func TestLeapYears(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := IsLeapYear(tc.year)
			if actual != tc.expected {
				t.Fatalf("IsLeapYear(%d) = %t, want %t", tc.year, actual, tc.expected)
			}
		})
	}
}

// Benchmark 400 years interval to get fair weighting of different years.
func Benchmark400(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for y := 1600; y < 2000; y++ {
			IsLeapYear(y)
		}
	}
}
