package dna

import "testing"

const targetTestVersion = 1

func (h Histogram) Equal(o Histogram) bool {
	return h.sameLength(o) && h.sameMappings(o)
}

func (h Histogram) sameLength(o Histogram) bool {
	return len(h) == len(o)
}

func (h Histogram) sameMappings(o Histogram) (res bool) {
	res = true
	for k := range h {
		if h[k] != o[k] {
			res = false
		}
	}
	return
}

func assertError(count, expected int, err error, t *testing.T) {
	var _ error = err
	if err == nil {
		t.Fatalf("Got \"%v\", expected \"%v\". error is nil", count, expected)
	}
}

func assertResult(count, expected int, err error, t *testing.T) {
	if err != nil {
		t.Fatalf("Raised error %v when expecting none", err)
	}
	if count != expected {
		t.Fatalf("Got \"%v\", expected \"%v\".", count, expected)
	}
}

var tallyTests = []struct {
	strand     string
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
		dna := DNA(tt.strand)
		count, err := dna.Count(tt.nucleotide)
		assertResult(count, tt.expected, err, t)
	}
}

func TestHasErrorForInvalidNucleotides(t *testing.T) {
	dna := DNA("GATTACA")
	count, err := dna.Count('X')
	assertError(count, 0, err, t)
}

// In most cases, this test is pointless.
// Very occasionally it matters.
// Just roll with it.
func TestCountingDoesntChangeCount(t *testing.T) {
	dna := DNA("CGATTGGG")
	dna.Count('T')
	count, err := dna.Count('T')
	assertResult(count, 2, err, t)
}

type histogramTest struct {
	strand   string
	expected Histogram
}

var histogramTests = []histogramTest{
	{
		"",
		Histogram{'A': 0, 'C': 0, 'T': 0, 'G': 0},
	},
	{
		"GGGGGGGG",
		Histogram{'A': 0, 'C': 0, 'T': 0, 'G': 8},
	},
	{
		"AGCTTTTCATTCTGACTGCAACGGGCAATATGTCTCTGTGTGGATTAAAAAAAGAGTGTCTGATAGCAGC",
		Histogram{'A': 20, 'C': 12, 'T': 21, 'G': 17},
	},
}

func TestSequenceHistograms(t *testing.T) {
	for _, tt := range histogramTests {
		dna := DNA(tt.strand)
		if !dna.Counts().Equal(tt.expected) {
			t.Fatalf("DNA{ \"%v\" }: Got \"%v\", expected \"%v\"", tt.strand, dna.Counts(), tt.expected)
		}
	}
}

func BenchmarkSequenceHistograms(b *testing.B) {
	b.StopTimer()
	for _, tt := range histogramTests {
		for i := 0; i < b.N; i++ {
			dna := DNA(tt.strand)
			b.StartTimer()

			dna.Counts()

			b.StopTimer()
		}
	}
}
