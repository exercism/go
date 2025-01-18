package variablelengthquantity

import (
	"bytes"
	"reflect"
	"testing"
)

func TestEncodeVarint(t *testing.T) {
	for _, tc := range encodeTestCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := EncodeVarint(tc.input); !bytes.Equal(actual, tc.expected) {
				t.Fatalf("EncodeVarint(%#v)\n got:%#v\nwant:%#v", tc.input, actual, tc.expected)
			}
		})
	}
}

func TestDecodeVarint(t *testing.T) {
	for _, tc := range decodeTestCases {
		t.Run(tc.description, func(t *testing.T) {
			actual, err := DecodeVarint(tc.input)
			switch {
			case tc.errorExpected:
				if err == nil {
					t.Fatalf("DecodeVarint(%#v) expected error, got: %#v", tc.input, actual)
				}
			case err != nil:
				t.Fatalf("DecodeVarint(%#v) returned error: %v, want:%#v", tc.input, err, tc.expected)
			case !reflect.DeepEqual(actual, tc.expected):
				t.Fatalf("DecodeVarint(%#v) = %#v, want:%#v", tc.input, actual, tc.expected)
			}
		})
	}
}
