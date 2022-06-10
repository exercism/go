package isbn

import (
	"testing"
)

func TestIsValidISBN(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			observed := IsValidISBN(test.isbn)
			if observed != test.expected {
				t.Errorf("IsValidISBN(%q)\nExpected: %t, Actual: %t",
					test.isbn, test.expected, observed)
			}
		})
	}
}

func BenchmarkIsValidISBN(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, n := range testCases {
			IsValidISBN(n.isbn)
		}
	}
}
