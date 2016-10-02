package variablelengthquantity

import (
	"bytes"
	"testing"
)

func TestEncodeDecodeVarint(t *testing.T) {
	testCases := []struct {
		input  []byte
		output uint32
		size   int
	}{
		0: {[]byte{0x7F}, 127, 1},
		1: {[]byte{0x81, 0x00}, 128, 1},
		2: {[]byte{0xC0, 0x00}, 8192, 1},
		3: {[]byte{0xFF, 0x7F}, 16383, 1},
		4: {[]byte{0x81, 0x80, 0x00}, 16384, 2},
		5: {[]byte{0xFF, 0xFF, 0x7F}, 2097151, 2},
		6: {[]byte{0x81, 0x80, 0x80, 0x00}, 2097152, 3},
		7: {[]byte{0xC0, 0x80, 0x80, 0x00}, 134217728, 3},
		8: {[]byte{0xFF, 0xFF, 0xFF, 0x7F}, 268435455, 3},

		9:  {[]byte{0x82, 0x00}, 256, 1},
		10: {[]byte{0x81, 0x10}, 144, 1},
	}

	for i, tc := range testCases {
		t.Logf("test case %d - %#v\n", i, tc.input)
		o, size := DecodeVarint(tc.input)
		if o != tc.output {
			t.Fatalf("expected %d\ngot\n%d\n", tc.output, o)
		}
		if size != tc.size {
			t.Fatalf("expected encoding size of %d bytes\ngot %d bytes\n", tc.size, size)
		}
		if encoded := EncodeVarint(tc.output); bytes.Compare(encoded, tc.input) != 0 {
			t.Fatalf("%d - expected %#v\ngot\n%#v\n", tc.output, tc.input, encoded)
		}
	}
}
