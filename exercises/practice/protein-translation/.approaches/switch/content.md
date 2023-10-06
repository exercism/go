# `switch` statement

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

This approach uses a [`switch`][switch] statement for looking up valid codons.
First, a couple of errors are defined.
Next, an empty `string` slice is defined.

The `FromCodon()` function is defined with [named return values][named-return-values] which are initialized with their [zero values][zero-values].
If the look-up of the codon is matched by any switch case except the codons for "STOP", then the protein is set.
If the codon look-up is "STOP", then the the error is set for a stopped codon.
If the look-up of the codon is not matched by any switch case, then the error is set for an invalid codon.
The function then returns whatever the values are for the protein and the error.


The `FromRNA()` function is also defined with [named return values][named-return-values] which are initialized with their [zero values][zero-values].
The [`for`][for] loop in `FromRNA()` starts at `0` and increments by the codon length of 3 until reaching the [len()][len] of the RNA strand.
In each iteration of the loop, a [slice][slice] of the RNA strand is passed to the `FromCodon()` function, which returns the protein and an error.
If the error is for a `STOP` codon, then the loop `break`s out of the loop.
If the error is for an invalid codon, then the function immediately returns an empty `string` slice and the error.
If the error is `nil`, then the protein is appended to the output slice of proteins.

When the loop successfully finishes, the function returns the slice of proteins and a `nil` error.

~~~~exercism/note
The `err` returned from the call to `FromCodon()` shadows the `err` named return value defined for `from RNA()`.
If an `err` is returned for a `STOP` codon, it falls out of scope when breaking from the loop,
and the `err` returned at the end of the function is the `err` named return value with the zero (`nil`) value.
More info on shadowed variables can be found [here](https://yourbasic.org/golang/gotcha-shadowing-variables/).
Because shadowing variables can be confusing, it may be preferred to use `return proteins, nil` for the final return statement.
~~~~

[switch]: https://go.dev/tour/flowcontrol/9
[named-return-values]: https://yourbasic.org/golang/named-return-values-parameters/
[zero-values]: https://yourbasic.org/golang/default-zero-value/
[for]: https://gobyexample.com/for
[len]: https://pkg.go.dev/builtin#len
[slice]: https://gobyexample.com/slices
