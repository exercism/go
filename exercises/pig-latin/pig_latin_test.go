package piglatin

import (
	"fmt"
	"testing"
)

func TestPigLatin(t *testing.T) {
	for _, test := range testCases {
		fmt.Printf("Test: %s\n", test.description)
		if pl := Sentence(test.input); pl != test.expected {
			t.Fatalf("FAIL: Sentence(%q) = %q, want %q.", test.input, pl, test.expected)
		}
	}
}
