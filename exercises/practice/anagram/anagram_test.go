package anagram

import (
	"slices"
	"testing"
)

func TestDetectAnagrams(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Detect(tc.subject, tc.candidates)
			slices.Sort(actual)
			slices.Sort(tc.expected)
			if !slices.Equal(tc.expected, actual) {
				t.Errorf("Detect(%q, %#v) = %#v, want: %#v", tc.subject, tc.candidates, actual, tc.expected)
			}
		})
	}
}

func BenchmarkDetectAnagrams(b *testing.B) {
	for range b.N {
		for _, tt := range testCases {
			Detect(tt.subject, tt.candidates)
		}
	}
}
