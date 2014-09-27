package hamming

import (
	"fmt"
	"testing"
)

const testVersion = 1

// Retired testVersions
// (none) df178dc24b57f7d4ccf54d47430192749d56898f

type testCase struct {
	expected         int
	expectedOk       bool
	strandA, strandB string
	description      string
}

var testCases = []testCase{
	{0, true, "", "", "no distance between empty strands"},
	{2, true, "AG", "CT", "complete hamming distance for small strands"},
	{0, true, "A", "A", "no distance between identical strands"},
	{1, true, "A", "G", "complete distance for single nucleotide strands"},
	{1, true, "AT", "CT", "small hamming distance"},
	{1, true, "GGACG", "GGTCG", "small hamming distance in longer strands"},
	{1, false, "AATG", "AAA", "disallow first strand longer"},
	{2, false, "ATA", "AGTG", "disallow second strand longer"},
	{4, true, "GATACA", "GCATAA", "large hamming distance"},
	{9, true, "GGACGGATTCTG", "AGGACGGATTCT", "hamming distance in very long strands"},
}

func TestHamming(t *testing.T) {
	if TestVersion != testVersion {
		t.Errorf("Found TestVersion = %v, want %v.", TestVersion, testVersion)
	}
	for _, tc := range testCases {
		switch observed, err := Distance(tc.strandA, tc.strandB); {
		case err != nil:
			if tc.expectedOk {
				fatal(t, tc, "Distance returned error: "+err.Error())
			}
		case !tc.expectedOk:
			fatal(t, tc, fmt.Sprintf("Distance returned %d.  Expected error.",
				observed))
		case observed != tc.expected:
			fatal(t, tc, fmt.Sprintf("Distance returned %d.  Want %d.",
				observed, tc.expected))
		}
	}
}

func fatal(t *testing.T, tc testCase, s string) {
	t.Fatalf(`%s:
%s
%s
%s`,
		tc.description,
		tc.strandA,
		tc.strandB,
		s,
	)
}

func BenchmarkHamming(b *testing.B) {
	// bench combined time to run through all test cases
	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			Distance(tc.strandA, tc.strandB)
		}
	}
}
