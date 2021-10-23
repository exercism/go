package dna

import (
	"errors"
	"strings"
)

// DNA is a list of nucleotides
type DNA string
type Histogram map[byte]int

const validNucleotides = "ACGT"

// Counts generates a histogram of valid nucleotides in given DNA.
// Returns error if DNA contains invalid nucleotide.
func (dna DNA) Counts() (Histogram, error) {
	var total int
	result := make(Histogram)

	for i := range validNucleotides {
		nucleotide := validNucleotides[i]
		result[nucleotide] = strings.Count(string(dna), string(nucleotide))
		total += result[nucleotide]
	}
	if total != len(dna) {
		return nil, errors.New("dna: contains invalid nucleotide")
	}
	return result, nil
}
