package dna

import (
	"errors"
	"strings"
)

// DNA is a list of nucleotides
type DNA string

var validNucleotides = []string{"A", "C", "G", "T"}

func isValidNucleotide(nucleotide string) bool {
	for _, n := range validNucleotides {
		if nucleotide == n {
			return true
		}
	}
	return false
}

// Count counts number of occurrences of given nucleotide in given DNA
func (dna DNA) Count(nucleotide string) (count int, err error) {
	if !isValidNucleotide(nucleotide) {
		return 0, errors.New("dna: invalid nucleotide " + string(nucleotide))
	}

	return strings.Count(string(dna), nucleotide), nil
}

// Counts generates a histogram of valid nucleotides in given DNA.
// Returns error if DNA contains invalid nucleotide.
func (dna DNA) Counts() (result map[string]int, e error) {
	var total int
	result = make(map[string]int)

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
