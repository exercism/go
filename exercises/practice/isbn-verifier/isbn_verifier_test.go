package isbnverifier

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
	for range b.N {
		for _, n := range testCases {
			IsValidISBN(n.isbn)
		}
	}
}
