# `switch` statement

```go
// Package dna is a small library for counting nucleotides in a DNA strand.
package dna

import "fmt"

// Histogram maps nucleotides to the count of their occurrence in a DNA strand.
type Histogram map[rune]int

// DNA is the series of nucleotides in the DNA strand.
type DNA string

// Counts generates a histogram of valid nucleotides for the given DNA strand.
func (dna DNA) Counts() (Histogram, error) {
	results := Histogram{
		'A': 0,
		'C': 0,
		'G': 0,
		'T': 0,
	}
	for _, nuc := range dna {
		switch nuc {
		case 'A', 'C', 'G', 'T':
			results[nuc]++
		default:
			return nil, fmt.Errorf("invalid nucleotide '%c'", nuc)
		}
	}
	return results, nil
}
```

In this approach a [`switch`][switch] statement is used for matching to a valid nucleotide.
The count in the histogram is incremented for the valid nucleotide.
If the character is not a valid nucleotide, then the function returns `nil` for the histogram and an error which identifies the invalid character.

If the iteration successfully finishes, then the function returns the histogram and a `nil` error.

[switch]: https://go.dev/tour/flowcontrol/9
