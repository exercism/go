package dna

import (
	"reflect"
	"testing"
)

var tallyTests = []struct {
	strand     DNA
	nucleotide byte
	expected   int
}{
	{"", 'A', 0},
	{"ACT", 'G', 0},
	{"CCCCC", 'C', 5},
	{"GGGGGTAACCCGG", 'T', 1},
}

func TestNucleotideCounts(t *testing.T) {
	for _, tt := range tallyTests {
		if count, err := tt.strand.Count(tt.nucleotide); err != nil {
			t.Fatal(err)
		} else if count != tt.expected {
			t.Fatalf("Got \"%v\", expected \"%v\"", count, tt.expected)
		}
	}
}

func TestHasErrorForInvalidNucleotides(t *testing.T) {
	dna := DNA("GATTACA")
	if _, err := dna.Count('X'); err == nil {
		t.Fatalf("X is an invalid nucleotide, but no error was raised")
	}
}

// In most cases, this test is pointless.
// Very occasionally it matters.
// Just roll with it.
func TestCountingDoesntChangeCount(t *testing.T) {
	dna := DNA("CGATTGGG")
	dna.Count('T')
	count1, err := dna.Count('T')
	if err != nil {
		t.Fatal(err)
	}
	count2, err := dna.Count('T')
	if err != nil {
		t.Fatal(err)
	}
	if count1 != count2 || count2 != 2 {
		t.Fatalf("Got %v, expected %v", []int{count1, count2}, []int{2, 2})
	}
}

type histogramTest struct {
	strand   DNA
	expected Histogram
	err      bool
}

var histogramTests = []histogramTest{
	{
		"",
		Histogram{'A': 0, 'C': 0, 'T': 0, 'G': 0},
		false,
	},
	{
		"GGGGGGGG",
		Histogram{'A': 0, 'C': 0, 'T': 0, 'G': 8},
		false,
	},
	{
		"AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC",
		Histogram{'A': 20, 'C': 12, 'T': 21, 'G': 17},
		false,
	},
	{
		"GGXXX",
		nil,
		true,
	},
}

func TestSequenceHistograms(t *testing.T) {
	for _, tt := range histogramTests {
		counts, err := tt.strand.Counts()
		if tt.err && err == nil {
			t.Fatalf("DNA{ %q }: expected error but didn't get one.", tt.strand)
		} else if !tt.err && err != nil {
			t.Fatalf("DNA{ %q }: expected no error but got error %s", tt.strand, err.Error())
		} else if !tt.err && !reflect.DeepEqual(counts, tt.expected) {
			t.Fatalf("DNA{ %q }: Got %v, expected %v", tt.strand, counts, tt.expected)
		}
	}
}

func BenchmarkSequenceHistograms(b *testing.B) {
	b.StopTimer()
	for _, tt := range histogramTests {
		for i := 0; i < b.N; i++ {
			b.StartTimer()

			tt.strand.Counts()

			b.StopTimer()
		}
	}
}

const targetTestVersion = 2

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Errorf("Found testVersion = %v, want %v.", testVersion, targetTestVersion)
	}
}
