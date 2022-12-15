# Iterate as runes

```go
// Package hamming is a small library for determining the hamming distance between two strands.
package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance accepts two strands and returns their hamming distance.
func Distance(strand1 string, strand2 string) (distance int, e error) {
	if utf8.RuneCountInString(strand1) != utf8.RuneCountInString(strand2) {
		return distance, errors.New("strands must be of equal length")
	}

	strand1Runes := []rune(strand1)
	strand2Runes := []rune(strand2)
	for idx, rooney := range strand1Runes {
		if rooney != strand2Runes[idx] {
			distance++
		}
	}
	return distance, e
}
```

The `Distance()` function is defined with [named return values][named-return-values] which are initialized with their [zero values][zero-values].

The [`utf8`][utf8] [`RuneCountInString()`][runecountinstring] function is used to compare the lengths of the strands in [runes][runes].
This isn't necessary when all of the input characters are single-byte [ASCII][ascii], but if any character was a multi-byte character,
then the `RuneCountInString()` function would be needed to give the accurate number of characters.
The [`len()`][len] function used on a string would give the number of [bytes][bytes],
which would be inaccurate for returning the number of characters for multi-byte characters.

If the strands don't have the same number of runes, then `distance` is returned (with its initial value of `0`) along with the new error.

The two strands are used to create rune slices.

The ways to iterate characters are by Unicode runes, or by each letter being a string, or by each letter being a byte.
The runes are from [`range`][range] on a string, the strings from [`Split()`][split], and the bytes from indexing into the string.
Another way to iterate runes is to convert the string to a rune [`slice`][slice] and `range` on it.
The difference between ranging on a rune slice vs ranging on a string is that the index returned from a string is the position of the next rune in bytes,
not which rune it is.
For example, if the first unicode character is two bytes,
then the second unicode character index will be 2 when ranging on a string and 1 when ranging on a rune slice.

The runes of the first strand are iterated.
The rune in each iteration is compared with the rune at the same index in the other strand.
If the runes are different at the same index for each strand, then the distance is incremented.
After the loop is finished, `distance` is returned along with `e` (with its initial value of `nil`).

[named-return-values]: https://yourbasic.org/golang/named-return-values-parameters/
[zero-values]: https://yourbasic.org/golang/default-zero-value/
[utf8]: https://pkg.go.dev/unicode/utf8
[runecountinstring]: https://pkg.go.dev/unicode/utf8#RuneCountInString
[runes]: https://golangdocs.com/rune-in-golang
[ascii]: https://www.asciitable.com/
[len]: https://pkg.go.dev/builtin#len
[bytes]: https://pkg.go.dev/bytes
[range]: https://gobyexample.com/range
[split]: https://pkg.go.dev/strings#Split
[slice]: https://gobyexample.com/slices
