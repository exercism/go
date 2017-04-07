package variablelengthquantity

import (
	"bytes"
	"testing"
)

const targetTestVersion = 2

var testCases = []struct {
	input  []byte
	output uint32
	size   int
}{
	0:  {[]byte{0x7F}, 127, 1},
	1:  {[]byte{0x81, 0x00}, 128, 2},
	2:  {[]byte{0xC0, 0x00}, 8192, 2},
	3:  {[]byte{0xFF, 0x7F}, 16383, 2},
	4:  {[]byte{0x81, 0x80, 0x00}, 16384, 3},
	5:  {[]byte{0xFF, 0xFF, 0x7F}, 2097151, 3},
	6:  {[]byte{0x81, 0x80, 0x80, 0x00}, 2097152, 4},
	7:  {[]byte{0xC0, 0x80, 0x80, 0x00}, 134217728, 4},
	8:  {[]byte{0xFF, 0xFF, 0xFF, 0x7F}, 268435455, 4},
	9:  {[]byte{0x82, 0x00}, 256, 2},
	10: {[]byte{0x81, 0x10}, 144, 2},
}

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}

func TestDecodeVarint(t *testing.T) {
	for i, tc := range testCases {
		o, size := DecodeVarint(tc.input)
		if o != tc.output {
			t.Fatalf("FAIL: case %d - expected %d got %d\n", i, tc.output, o)
		} else if size != tc.size {
			t.Fatalf("FAIL: case %d - expected encoding size of %d bytes\ngot %d bytes\n", i, tc.size, size)
		} else {
			t.Logf("PASS: case %d - %#v\n", i, tc.input)
		}
	}
}

func TestEncodeVarint(t *testing.T) {
	for i, tc := range testCases {
		if encoded := EncodeVarint(tc.output); bytes.Compare(encoded, tc.input) != 0 {
			t.Fatalf("FAIL: case %d - %d \nexpected\t%#v\ngot\t\t%#v\n", i, tc.output, tc.input, encoded)
		} else {
			t.Logf("PASS: case %d - %#v\n", i, tc.input)
		}
	}
}
