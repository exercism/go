package dna

import (
	"errors"
	"strings"
)

const testVersion = 2

// Histogram is a mapping from nucleotide to its count in given DNA
type Histogram map[byte]int

// DNA is a list of nucleotides
type DNA string

const validNucleotides = "ACGT"

// Count counts number of occurrences of given nucleotide in given DNA
func (dna DNA) Count(nucleotide byte) (count int, err error) {
	if !strings.Contains(validNucleotides, string(nucleotide)) {
		return 0, errors.New("dna: invalid nucleotide " + string(nucleotide))
	}

	return strings.Count(string(dna), string(nucleotide)), nil
}

// Counts generates a histogram of valid nucleotides in given DNA.
// Returns error if DNA contains invalid nucleotide.
func (dna DNA) Counts() (Histogram, error) {
	var total int
	h := Histogram{}
	for i := range validNucleotides {
		nucleotide := validNucleotides[i]
		h[nucleotide], _ = dna.Count(nucleotide)
		total += h[nucleotide]
	}
	if total != len(dna) {
		return nil, errors.New("dna: contains invalid nucleotide")
	}
	return h, nil
}
