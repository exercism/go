# Introduction

There are at least two idomatic approaches to solving Protein Translation in Go.
You can use the a `switch` statement to lookup up the RNA.
Or you can use a `map`.

## General guidance

The key to solving Protein Translation is to look up valid RNA and to error on invalid codons.

## Approach: `switch` statement

```go
// Package protein has functionality to translate RNA sequences into proteins.
package protein

import (
	"errors"
)

var (
	// ErrStop indicates the translation was stopped.
	ErrStop = errors.New("stop")

	// ErrInvalidBase indicates the codon was invalid.
	ErrInvalidBase = errors.New("invalid codon")

	noProteins = []string{}
)

const codonLength = 3

// FromCodon either translates a codon to a protein or returns an error for a stop codon or invalid codon.
func FromCodon(codon string) (protein string, err error) {
	switch codon {
	case "AUG":
		protein = "Methionine"
	case "UUU", "UUC":
		protein = "Phenylalanine"
	case "UUA", "UUG":
		protein = "Leucine"
	case "UCU", "UCC", "UCA", "UCG":
		protein = "Serine"
	case "UAU", "UAC":
		protein = "Tyrosine"
	case "UGU", "UGC":
		protein = "Cysteine"
	case "UGG":
		protein = "Tryptophan"
	case "UAA", "UAG", "UGA":
		err = ErrStop
	default:
		err = ErrInvalidBase
	}

	return protein, err
}

// FromRNA maps RNA codons to their matching proteins and returns either the list of proteins
// or returns an error for an invalid codon.
func FromRNA(rna string) (proteins []string, err error) {
	for i := 0; i < len(rna); i += codonLength {
		protein, err := FromCodon(rna[i : i+codonLength])
		if err == ErrStop {
			break
		}
		if err == ErrInvalidBase {
			return noProteins, err
		}
		proteins = append(proteins, protein)
	}

	return proteins, err
}
```

For more information, check the [`switch` statement approach][approach-switch-statement].

## Approach: `map`

```go
// Package protein has functionality to translate RNA sequences into proteins.
package protein

import "errors"

var (
	// ErrInvalidBase indicates the codon was invalid.
	ErrInvalidBase = errors.New("invalid codon")
	// ErrStop indicates the translation was stopped.
	ErrStop = errors.New("stop")

	noProteins = []string{}
	lookup = map[string]string{
		"AUG": "Methionine",
		"UUU": "Phenylalanine", "UUC": "Phenylalanine",
		"UUA": "Leucine", "UUG": "Leucine",
		"UCU": "Serine", "UCC": "Serine", "UCA": "Serine", "UCG": "Serine",
		"UAU": "Tyrosine", "UAC": "Tyrosine",
		"UGU": "Cysteine", "UGC": "Cysteine",
		"UGG": "Tryptophan",
		"UAA": "STOP", "UAG": "STOP", "UGA": "STOP",
	}
)

const codonLength = 3

// FromRNA maps RNA codons to their matching proteins and returns either the list of proteins
// or returns an error for an invalid codon.
func FromRNA(rna string) (proteins []string, err error) {
	for i := 0; i < len(rna); i += codonLength {
		codon, err := FromCodon(rna[i : i+codonLength])
		if err == ErrStop {
			break
		}
		if err == ErrInvalidBase {
			return noProteins, err
		}
		proteins = append(proteins, codon)
	}
	return proteins, err
}

// FromCodon either translates a codon to a protein or returns an error for a stop codon or invalid codon.
func FromCodon(codon string) (protein string, err error) {
	protein = lookup[codon]
	if protein == "" {
		err = ErrInvalidBase
	} else if protein == "STOP" {
		protein = ""
		err = ErrStop
	}
	return protein, err
}
```

For more information, check the [`map` approach][approach-map].

## Which approach to use?

Both approaches are about the same performance, so which to use could be a matter of preferred style.
To compare performance of the approaches, check the [Performance article][article-performance].

[approach-switch-statement]: https://exercism.org/tracks/go/exercises/protein-translation/approaches/switch
[approach-map]: https://exercism.org/tracks/go/exercises/protein-translation/approaches/map
[article-performance]: https://exercism.org/tracks/go/exercises/protein-translation/articles/performance
