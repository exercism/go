package isbn

import (
	"testing"
)

func TestIsValidISBN(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := IsValidISBN(tc.isbn)
			if actual != tc.expected {
				t.Errorf("IsValidISBN(%q)=%t, want: %t", tc.isbn, actual, tc.expected)
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
