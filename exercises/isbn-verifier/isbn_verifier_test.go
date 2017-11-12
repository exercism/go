package isbn

import (
	"testing"
)

func TestIsValidISBN(t *testing.T) {
	for _, test := range testCases {
		observed := IsValidISBN(test.isbn)
		if observed != test.expected {
			t.Errorf("FAIL: %s\nIsValidISBN(%q)\nExpected: %t, Actual: %t",
				test.description, test.isbn, test.expected, observed)
		}
		t.Logf("PASS: %s", test.description)
	}
}
