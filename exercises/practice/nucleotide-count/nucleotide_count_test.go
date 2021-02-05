package dna

import (
	"reflect"
	"testing"
)

func TestCounts(t *testing.T) {
	for _, tc := range testCases {
		dna := DNA(tc.strand)
		s, err := dna.Counts()
		switch {
		case tc.errorExpected:
			if err == nil {
				t.Fatalf("FAIL: %s\nCounts(%q)\nExpected error\nActual: %#v",
					tc.description, tc.strand, s)
			}
		case err != nil:
			t.Fatalf("FAIL: %s\nCounts(%q)\nExpected: %#v\nGot error: %q",
				tc.description, tc.strand, tc.expected, err)
		case !reflect.DeepEqual(s, tc.expected):
			t.Fatalf("FAIL: %s\nCounts(%q)\nExpected: %#v\nActual: %#v",
				tc.description, tc.strand, tc.expected, s)
		}
		t.Logf("PASS: %s", tc.description)
	}
}
