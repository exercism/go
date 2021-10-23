package octal

import (
	"testing"
)

var testCases = []struct {
	input       string
	expectedNum int64
	expectErr   bool
}{
	{"1", 1, false},
	{"10", 8, false},
	{"1234567", 342391, false},
	{"carrot", 0, true},
	{"35682", 0, true},
}

func TestParseOctal(t *testing.T) {
	for _, test := range testCases {
		actualNum, actualErr := ParseOctal(test.input)
		// check actualNum only if no error expected
		if !test.expectErr && actualNum != test.expectedNum {
			t.Fatalf("ParseOctal(%s): expected[%d], actual [%d]",
				test.input, test.expectedNum, actualNum)
		}
		// if we expect an error and there isn't one
		if test.expectErr && actualErr == nil {
			t.Errorf("ParseOctal(%s): expected an error, but error is nil", test.input)
		}
		// if we don't expect an error and there is one
		if !test.expectErr && actualErr != nil {
			var _ error = actualErr
			t.Errorf("ParseOctal(%s): expected no error, but error is: %s", test.input, actualErr)
		}
	}
}

func BenchmarkParseOctal(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}

	for i := 0; i < b.N; i++ {

		for _, test := range testCases {
			ParseOctal(test.input)
		}

	}
}
