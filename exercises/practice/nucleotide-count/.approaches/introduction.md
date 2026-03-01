# Introduction

There are at least two idomatic approaches to solving Nucleotide Count in Go.
You can use the a `switch` statement to lookup up the DNA.
Or you can use `if` with a short statement.

## General guidance

The key to solving Nucleotide Count is to look up valid nucleotides and to error on invalid input.

## Approach: `switch` statement

```go
// Package dna is a small library for counting nucleotides in a DNA strand.
package nucleotidecount

import "fmt"

// Histogram maps nucleotides to the count of their occurrence in a DNA strand.
type Histogram map[rune]int

// DNA is the series of nucleotides in the DNA strand.
type DNA string

// Counts generates a histogram of valid nucleotides for the given DNA strand.
func (dna DNA) Counts() (Histogram, error) {
	results := Histogram{'A': 0, 'C': 0, 'G': 0, 'T': 0}
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

For more information, check the [`switch` statement approach][approach-switch-statement].

## Approach: `if` with a short statement

```go
// Package dna is a small library for counting nucleotides in a DNA strand.
package nucleotidecount

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

For more information, check the [`if` with short statement approach][approach-if-with-short-statement].

## Which approach to use?

The `switch` statement approach benchmarked the fastest.
To compare performance of the approaches, check the [Performance article][article-performance].

[approach-switch-statement]: https://exercism.org/tracks/go/exercises/nucleotide-count/approaches/switch-statement
[approach-if-with-short-statement]: https://exercism.org/tracks/go/exercises/nucleotide-count/approaches/if-with-short-statement
[article-performance]: https://exercism.org/tracks/go/exercises/nucleotide-count/articles/performance
