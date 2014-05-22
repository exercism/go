package hexadecimal

import (
	"fmt"
	"reflect"
	"testing"
)

var testCases = []struct {
	input           string
	expectedDecimal int64
	expectedErr     error
}{
	{"1", 1, nil},
	{"10", 16, nil},
	{"2d", 45, nil},
	{"cfcfcf", 13619151, nil},
	{"CFCFCF", 13619151, nil},
	{"peanut", 0, fmt.Errorf("peanut is an invalid hexadecimal string")},
	{"2cg134", 0, fmt.Errorf("2cg134 is an invalid hexadecimal string")},
}

func TestToDecimal(t *testing.T) {
	for _, test := range testCases {
		actual, err := ToDecimal(test.input)
		if actual != test.expectedDecimal {
			t.Fatalf("ToDeciaml(%s): expected result[%d], actual[%d]", test.input, test.expectedDecimal, actual)
		}
		if !reflect.DeepEqual(err, test.expectedErr) {
			t.Fatalf("ToDeciaml(%s): expected error[%s], actual[%s]", test.input, test.expectedErr, err)
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
