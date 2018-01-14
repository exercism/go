package dna

import (
	"errors"
	"strings"
)

// DNA is a list of nucleotides
type DNA string
type Histogram map[byte]int

const validNucleotides = "ACGT"

func isValidNucleotide(nucleotide byte) bool {
	return strings.Contains(validNucleotides, string(nucleotide))
}

// Count counts number of occurrences of given nucleotide in given DNA
func (dna DNA) Count(nucleotide byte) (count int, err error) {
	if !isValidNucleotide(nucleotide) {
		return 0, errors.New("dna: invalid nucleotide " + string(nucleotide))
	}
	return strings.Count(string(dna), string(nucleotide)), nil
}

// Counts generates a histogram of valid nucleotides in given DNA.
// Returns error if DNA contains invalid nucleotide.
func (dna DNA) Counts() (result Histogram, e error) {
	var total int
	result = make(Histogram)

	for i := range validNucleotides {
		nucleotide := validNucleotides[i]
		result[nucleotide], _ = dna.Count(nucleotide)
		total += result[nucleotide]
	}
	if total != len(dna) {
		return nil, errors.New("dna: contains invalid nucleotide")
	}
	return result, nil
}
