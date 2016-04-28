package variablelengthquantity

import (
	"bytes"
	"testing"
)

/*
The goal of this exercise is to implement
VLQ encoding/decoding as described here: https://en.wikipedia.org/wiki/Variable-length_quantity

In short, the goal of this encoding is to save encode integer values in a way that would save bytes.
Only the first 7 bits of each byte is significant (right-justified; sort of like an ASCII byte).
So, if you have a 32-bit value, you have to unpack it into a series of 7-bit bytes.
Of course, you will have a variable number of bytes depending upon your integer.
To indicate which is the last byte of the series, you leave bit #7 clear.
In all of the preceding bytes, you set bit #7.

So, if an integer is between `0-127`, it can be represented as one byte.
The largest integer allowed is `0FFFFFFF`, which translates to 4 bytes variable length.
Here are examples of delta-times as 32-bit values, and the variable length quantities that they translate to:
*/

func TestEncodeDecodeVarint(t *testing.T) {
	testCases := []struct {
		input  []byte
		output uint32
	}{
		0: {[]byte{0x7F}, 127},
		1: {[]byte{0x81, 0x00}, 128},
		2: {[]byte{0xC0, 0x00}, 8192},
		3: {[]byte{0xFF, 0x7F}, 16383},
		4: {[]byte{0x81, 0x80, 0x00}, 16384},
		5: {[]byte{0xFF, 0xFF, 0x7F}, 2097151},
		6: {[]byte{0x81, 0x80, 0x80, 0x00}, 2097152},
		7: {[]byte{0xC0, 0x80, 0x80, 0x00}, 134217728},
		8: {[]byte{0xFF, 0xFF, 0xFF, 0x7F}, 268435455},

		9:  {[]byte{0x82, 0x00}, 256},
		10: {[]byte{0x81, 0x10}, 144},
	}

	for i, tc := range testCases {
		t.Logf("test case %d - %#v\n", i, tc.input)
		if o, _ := DecodeVarint(tc.input); o != tc.output {
			t.Fatalf("expected %d\ngot\n%d\n", tc.output, o)
		}
		if encoded := EncodeVarint(tc.output); bytes.Compare(encoded, tc.input) != 0 {
			t.Fatalf("%d - expected %#v\ngot\n%#v\n", tc.output, tc.input, encoded)
		}
	}
}
