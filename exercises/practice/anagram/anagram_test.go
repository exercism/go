package anagram

import (
	"fmt"
	"sort"
	"testing"
)

func equal(a, b []string) bool {
	if len(b) != len(a) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func TestDetectAnagrams(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Detect(tc.subject, tc.candidates)
			if !equal(tc.expected, actual) {
				t.Errorf("Detect(%q, %#v) = %#v, want: %#v", tc.subject, tc.candidates, actual, tc.expected)
			}
		})
	}
}

func BenchmarkDetectAnagrams(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	for i := 0; i < b.N; i++ {
		for _, tt := range testCases {
			Detect(tt.subject, tt.candidates)
		}
	}
}
