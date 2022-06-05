package piglatin

import (
	"testing"
)

func TestPigLatin(t *testing.T) {
	for _, test := range testCases {
		t.Run(test.description, func(t *testing.T) {
			if pl := Sentence(test.input); pl != test.expected {
				t.Fatalf("Sentence(%q) = %q, want %q.", test.input, pl, test.expected)
			}
		})
	}
}
