package dna

import (
	"errors"
	"strings"
)

const testVersion = 1

// Histogram is a mapping from nucleotide to its count in given DNA
type Histogram map[byte]int

// DNA is a list of nucleotides
type DNA string

// Count counts number of occurrences of given nucleotide in given DNA
func (dna DNA) Count(nucleotide byte) (count int, err error) {
	validNucleotides := "ACGT"
	if !strings.Contains(validNucleotides, string(nucleotide)) {
		return 0, errors.New("dna: invalid nucleotide " + string(nucleotide))
	}

	return strings.Count(string(dna), string(nucleotide)), nil
}

// Counts generates a histogram of valid nucleotides in given DNA.
// Returns error if DNA contains invalid nucleotide.
func (dna DNA) Counts() Histogram {
	a, _ := dna.Count('A')
	c, _ := dna.Count('C')
	t, _ := dna.Count('T')
	g, _ := dna.Count('G')

	return Histogram{'A': a, 'C': c, 'T': t, 'G': g}
}
