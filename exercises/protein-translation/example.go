// Package protein translates RNA sequences into proteins.
package protein

import "errors"

var (
	STOP           = errors.New("stop")
	ErrInvalidBase = errors.New("invalid base")
)

// IsStopCodon checks whether a codon is a stop codon.
func IsStopCodon(c string) bool {
	return c == "UAA" || c == "UAG" || c == "UGA"

}

// FromCodon returns the protein for the given codon.
// If the codon is a stop codon, it returns STOP.
// If the codon is unrecognized, it returns invalidBase.
func FromCodon(c string) (string, error) {
	if IsStopCodon(c) {
		return "", STOP
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
func FromRNA(s string) []string {
	var proteins []string
	bases := []rune(s)
	for i := 0; i < len(bases); i += 3 {
		p, err := FromCodon(string(bases[i : i+3]))
		if err == STOP {
			break
		}
		proteins = append(proteins, p)
	}
	return proteins
}
