package dna

import (
	"reflect"
	"testing"
)

func TestCounts(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			dna := DNA(tc.strand)
			actual, err := dna.Counts()
			switch {
			case tc.errorExpected:
				if err == nil {
					t.Fatalf("DNA.Counts(%q) expected error, got: %#v", tc.strand, actual)
				}
			case err != nil:
				t.Fatalf("DNA.Counts(%q) returned error: %v, want: %#v", tc.strand, err, tc.expected)
			case !reflect.DeepEqual(actual, tc.expected):
				t.Fatalf("DNA.Counts(%q)\n got:%#v\nwant: %#v", tc.strand, actual, tc.expected)
			}
		})
	}
}
