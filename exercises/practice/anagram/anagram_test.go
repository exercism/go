package anagram

import (
	"fmt"
	"sort"
	"testing"
)

var nonAsciiTestCases = []anagramTest{
	{
		description: "detects unicode anagrams",
		subject:     "ΑΒΓ",
		candidates:  []string{"ΒΓΑ", "ΒΓΔ", "γβα"},
		expected:    []string{"ΒΓΑ", "γβα"},
	},
}

func TestDetectAnagrams(t *testing.T) {
	var allCases = append(testCases, nonAsciiTestCases...)
	for _, tc := range allCases {
		t.Run(tc.description, func(t *testing.T) {
			actual := Detect(tc.subject, tc.candidates)
			if !equal(tc.expected, actual) {
				t.Errorf("Detect(%q, %#v) = %#v, want: %#v", tc.subject, tc.candidates, actual, tc.expected)
			}
		})
	}
}

func equal(a, b []string) bool {
	if len(b) != len(a) {
		return false
	}

	sort.Strings(a)
	sort.Strings(b)
	return fmt.Sprintf("%v", a) == fmt.Sprintf("%v", b)
}

func BenchmarkDetectAnagrams(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark in short mode.")
	}
	var allCases = append(testCases, nonAsciiTestCases...)
	for i := 0; i < b.N; i++ {
		for _, tt := range allCases {
			Detect(tt.subject, tt.candidates)
		}
	}
}
