package igpay

import "testing"

func TestPigLatin(t *testing.T) {
	for _, test := range testCases {
		if pl := PigLatin(test.input); pl != test.expected {
      t.Fatalf("FAIL: %s\nPigLatin(%q) = %q, want %q.", test.description, test.input, pl, test.expected)
		}
	}
}
