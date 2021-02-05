package isbn

import (
	"testing"
)

func TestIsValidISBN(t *testing.T) {
	for _, test := range testCases {
		observed := IsValidISBN(test.isbn)
		if observed == test.expected {
			t.Logf("PASS: %s", test.description)
		} else {
			t.Errorf("FAIL: %s\nIsValidISBN(%q)\nExpected: %t, Actual: %t",
				test.description, test.isbn, test.expected, observed)
		}
	}
}

func BenchmarkIsValidISBN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, n := range testCases {
			IsValidISBN(n.isbn)
		}
	}
}
