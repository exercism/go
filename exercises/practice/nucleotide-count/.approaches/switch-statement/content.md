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

In this approach a [`switch`][switch] statement is used for matching to a valid nucleotide.
The count in the histogram is incremented for the valid nucleotide.
If the character is not a valid nucleotide, then the function returns `nil` for the histogram and an error which identifies the invalid character.

If the iteration successfully finishes, then the function returns the histogram and a `nil` error.

# Iterate by bytes

Since iterating by bytes is often more performant than iterating by runes, it may be of interest to try that for Nucleotide Count.

The ways to iterate characters are by Unicode runes, or by each letter being a string, or by each letter being a byte.
The runes are from [`range`][range] on a string, the strings from [`Split()`][split], and the bytes from indexing into the string.
Another way to iterate runes is to convert the string to a rune [`slice`][slice] and `range` on it.
The difference between ranging on a rune slice vs ranging on a string is that the index returned from a string is the position of the next rune in bytes,
not which rune it is.
For example, if the first unicode character is two bytes,
then the second unicode character index will be 2 when ranging on a string and 1 when ranging on a rune slice.
As of the time of this writing we can iterate bytes, since all of the characters are ASCII.

You can modify the code to look something like this

```go
// Package dna is a small library for counting nucleotides in a DNA strand.
package dna

import "fmt"

// Histogram maps nucleotides to the count of their occurrence in a DNA strand.
type Histogram map[byte]int

// DNA is the series of nucleotides in the DNA strand.
type DNA string

const (
	nucA byte = 65
	nucC byte = 67
	nucG byte = 71
	nucT byte = 84
)

// Counts generates a histogram of valid nucleotides for the given DNA strand.
func (dna DNA) Counts() (Histogram, error) {
	results := Histogram{nucA: 0, nucC: 0, nucG: 0, nucT: 0}
	length := len(dna)
	for i := 0; i < length; i++ {
		nuc := dna[i]
		switch nuc {
		case nucA, nucC, nucG, nucT:
			results[nuc]++
		default:
			return nil, fmt.Errorf("invalid nucleotide '%c'", nuc)
		}
	}
	return results, nil
}
```


The Histogram is defined as `map[byte]int` instead of `map[rune]int`.
[Constants][const] are defined with the [ASCII][ascii] values of the valid DNA characters.

The [`len()`][len] function is used to get the length of the DNA strand.
The `for` loop iterates from `0` until the length of the strand.
Indexing into a string gets the byte at that position.

As when iterating by runes, a [`switch`][switch] statement is used for matching to a valid nucleotide.
The count in the histogram is incremented for the valid nucleotide.
If the character is not a valid nucleotide, then the function returns `nil` for the histogram and an error which identifies the invalid character.

If the iteration successfully finishes, then the function returns the histogram and a `nil` error.
However, the tests expect the Histogram to be returned as a `map[rune]int`.
An ASCII byte can fit in a rune, but that takes time, as seen by the benchmark

```
switch on runes
BenchmarkCounts-12    	  593500	      1942 ns/op	     840 B/op	      12 allocs/o

switch on bytes
BenchmarkCounts-12    	  441386	      2645 ns/op	     760 B/op	      12 allocs/op
```

Since a `map[rune]int` is expected to be returned, it is faster to iterate by runes in this situation.

[switch]: https://go.dev/tour/flowcontrol/9
[ascii]: https://www.asciitable.com/
[range]: https://gobyexample.com/range
[split]: https://pkg.go.dev/strings#Split
[slice]: https://gobyexample.com/slices
[const]: https://go.dev/tour/basics/15
[len]: https://pkg.go.dev/builtin#len
