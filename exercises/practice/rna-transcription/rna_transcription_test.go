package rnatranscription

import "testing"

func TestRNATranscription(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			if actual := ToRNA(tc.input); actual != tc.expected {
				t.Fatalf("ToRNA(%q) = %q, want: %q", tc.input, actual, tc.expected)
			}
		})
	}
}

func BenchmarkRNATranscription(b *testing.B) {
	for range b.N {
		for _, tc := range testCases {
			ToRNA(tc.input)
		}
	}
}
