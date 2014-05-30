package hexadecimal

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	in  string
	out int64
	err error
}{
	{"1", 1, nil},
	{"10", 0x10, nil},
	{"2d", 0x2d, nil},
	{"012", 0x12, nil},
	{"cfcfcf", 0xcfcfcf, nil},
	{"CFCFCF", 0xcfcfcf, nil},
	{"", 0, ErrSyntax},
	{"peanut", 0, ErrSyntax},
	{"2cg134", 0, ErrSyntax},
	{"8000000000000000", 1<<63 - 1, ErrRange},
	{"9223372036854775809", 1<<63 - 1, ErrRange},
}

// Modify the test cases to expect a &ParseError instead of
// the listed ErrSyntax or ErrRange.
func init() {
	for i := range testCases {
		test := &testCases[i]
		if test.err != nil {
			test.err = &ParseError{test.in, test.err}
		}
	}
}

func TestParseHex(t *testing.T) {
	for _, test := range testCases {
		out, err := ParseHex(test.in)
		if test.out != out || !reflect.DeepEqual(test.err, err) {
			t.Errorf("ParseHex(%q) = %v, %v want %v, %v",
				test.in, out, err, test.out, test.err)
		}
	}
}

func BenchmarkParseHex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			ParseHex(test.in)
		}
	}
}
