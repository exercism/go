// Package protein translates RNA sequences into proteins.
package protein

import (
	"errors"
)

// ErrStop indicates that one of the stop codons was read.
var ErrStop = errors.New("stop")

// ErrInvalidBase indicates that an unrecognized codon was read.
var ErrInvalidBase = errors.New("invalid base")

// IsStopCodon checks whether a codon is a stop codon.
func IsStopCodon(c string) bool {
	return c == "UAA" || c == "UAG" || c == "UGA"

}

// FromCodon returns the protein for the given codon.
// If the codon is a stop codon, it returns ErrStop.
// If the codon is unrecognized, it returns ErrInvalidBase.
func FromCodon(c string) (string, error) {
	if IsStopCodon(c) {
		return "", ErrStop
	}

	switch c {
	case
		"AUG":
		return "Methionine", nil
	case
		"UUU",
		"UUC":
		return "Phenylalanine", nil
	case
		"UUA",
		"UUG":
		return "Leucine", nil
	case
		"UCU",
		"UCC",
		"UCA",
		"UCG":
		return "Serine", nil
	case
		"UAU",
		"UAC":
		return "Tyrosine", nil
	case
		"UGU",
		"UGC":
		return "Cysteine", nil
	case
		"UGG":
		return "Tryptophan", nil
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA returns the protein sequence for an RNA strand.
// If the sequence contains invalid characters, it returns ErrInvalidBase.
func FromRNA(s string) ([]string, error) {
	var proteins []string
	bases := []rune(s)
	for i := 0; i < len(bases); i += 3 {
		p, err := FromCodon(string(bases[i : i+3]))
		switch err {
		case ErrStop:
			return proteins, nil
		case ErrInvalidBase:
			return nil, err
		default:
			proteins = append(proteins, p)
		}
	}
	return proteins, nil
}
