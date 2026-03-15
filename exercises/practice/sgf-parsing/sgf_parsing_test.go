package sgfparsing

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got, err := Parse(tc.encoded)
			if tc.expectedError != "" {
				if err == nil {
					t.Fatalf("Parse(%q) expected error, got: %v", tc.encoded, got)
				}
			} else if err != nil {
				t.Fatalf("Parse(%q) returned error: %q", tc.encoded, err)
			} else if !reflect.DeepEqual(got, tc.expected) {
				t.Fatalf("Parse(%q)\ngot:  %+v\nwant: %+v", tc.encoded, got, tc.expected)
			}
		})
	}
}

func BenchmarkParse(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			Parse(tc.encoded)
		}
	}
}
