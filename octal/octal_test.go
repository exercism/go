package octal

import (
	"testing"
)

var testCases = []struct {
	input           string
	expectedDecimal int64
	expectErr       bool
}{
	{"1", 1, false},
	{"10", 8, false},
	{"1234567", 342391, false},
	{"carrot", 0, true},
	{"35682", 0, true},
}

func TestToDecimal(t *testing.T) {
	for _, test := range testCases {
		actualDecimal, actualErr := ToDecimal(test.input)
		if actualDecimal != test.expectedDecimal {
			t.Fatalf("ToDecimal(%s): expected[%d], actual [%d]",
				test.input, test.expectedDecimal, actualDecimal)
		}

		// if we expect an error and there isn't one
		if test.expectErr && actualErr == nil {
			t.Errorf("ToDecimal(%s): expected an error, but error is nil", test.input)
		}
		// if we don't expect an error and there is one
		if !test.expectErr && actualErr != nil {
			t.Errorf("ToDecimal(%s): expected no error, but error is: %s", test.input, actualErr)
		}
	}
}

func BenchmarkToDecimal(b *testing.B) {
	b.StopTimer()

	for _, test := range testCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			ToDecimal(test.input)
		}

		b.StopTimer()
	}
}
