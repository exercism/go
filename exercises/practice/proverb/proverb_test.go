package proverb

import (
	"fmt"
	"testing"
)

func TestProverb(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Proverb(tc.input)
			if fmt.Sprintf("%q", actual) != fmt.Sprintf("%q", tc.expected) {
				t.Fatalf("Proverb(%#v)\n got:%#v\nwant:%#v", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkProverb(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, test := range testCases {
			Proverb(test.input)
		}
	}
}
