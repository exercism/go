# `if` with a short statement

```go
// Package dna is a small library for counting nucleotides in a DNA strand.
package dna

import (
	"fmt"
)

// Histogram maps nucleotides to the count of their occurrence in a DNA strand.
type Histogram = map[rune]int

// DNA is the series of nucleotides in the DNA strand..
type DNA string

// Counts generates a histogram of valid nucleotides for the given DNA strand.
func (dna DNA) Counts() (Histogram, error) {
	results := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
	for _, nuc := range dna {
		if _, ok := results[nuc]; ok {
			results[nuc]++
		} else {
			return nil, fmt.Errorf("invalid nucleotide '%c'", nuc)
		}
	}
	return results, nil
}
```

In this approach an [`if` with a short statement][if-with-short-statement] is used for looking up the character in the histogram.
If the character is a valid nucleotide, then the count in the histogram is incremented for it.
If the character is not a valid nucleotide, then the function returns `nil` for the histogram and an error which identifies the invalid character.

If the iteration successfully finishes, then the function returns the histogram and a `nil` error.

[if-with-short-statement]: https://go.dev/tour/flowcontrol/6
