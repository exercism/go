package piglatin

import (
	"fmt"
	"testing"
)

func TestPigLatin(t *testing.T) {
	for _, test := range testCases {
		if pl := Sentence(test.input); pl != test.expected {
			t.Fatalf("FAIL: Sentence(%q) = %q, want %q.", test.input, pl, test.expected)
		}
		fmt.Printf("PASS: %s\n", test.description)
	}
}
