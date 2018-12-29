package variablelengthquantity

import (
	"bytes"
	"reflect"
	"testing"
)

func TestDecodeVarint(t *testing.T) {
	for i, tc := range decodeTestCases {
		o, err := DecodeVarint(tc.input)
		if err != nil {
			if !tc.shouldFail {
				t.Fatalf("FAIL: case %d | %s\nexpected %#v got error: %q\n", i, tc.description, tc.output, err)
			}
		} else if tc.output == nil {
			t.Fatalf("FAIL: case %d | %s\nexpected error, got %#v\n", i, tc.description, o)
		} else if !reflect.DeepEqual(o, tc.output) {
			t.Fatalf("FAIL: case %d | %s\nexpected\t%#v\ngot\t\t%#v\n", i, tc.description, tc.output, o)
		}
		t.Logf("PASS: case %d | %s\n", i, tc.description)
	}
}

func TestEncodeVarint(t *testing.T) {
	for i, tc := range encodeTestCases {
		if encoded := EncodeVarint(tc.input); bytes.Compare(encoded, tc.output) != 0 {
			t.Fatalf("FAIL: case %d | %s\nexpected\t%#v\ngot\t\t%#v\n", i, tc.description, tc.output, encoded)
		}
		t.Logf("PASS: case %d | %s\n", i, tc.description)
	}
}
