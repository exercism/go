package hamming

import (
	"testing"
)

var testCases = []struct {
	expected         int
	strandA, strandB string
	description      string
}{
	{0, "", "", "no difference between empty strands"},
	{2, "AG", "CT", "complete hamming distance for small strands"},
	{0, "A", "A", "no difference between identical strands"},
	{1, "A", "G", "complete distance for single nucleotide strands"},
	{1, "AT", "CT", "small hamming distance"},
	{1, "GGACG", "GGTCG", "small hamming distance in longer strands"},
	{1, "AATG", "AAA", "ignores extra length on first strand when longer"},
	{2, "ATA", "AGTG", "ignores extra length on second strand when longer"},
	{4, "GATACA", "GCATAA", "large hamming distance"},
	{9, "GGACGGATTCTG", "AGGACGGATTCT", "hamming distance in very long strands"},
}

func TestHamming(t *testing.T) {
	for _, tc := range testCases {

		observed := Distance(tc.strandA, tc.strandB)

		if tc.expected != observed {
			t.Fatalf(`%s:
{%v,%v}
expected: %v
observed: %v`,
				tc.description,
				tc.strandA,
				tc.strandB,
				tc.expected,
				observed,
			)
		}
	}
}

func BenchmarkHamming(b *testing.B) {
	b.StopTimer()

	for _, tc := range testCases {
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			Distance(tc.strandA, tc.strandB)
		}

		b.StopTimer()
	}
}
