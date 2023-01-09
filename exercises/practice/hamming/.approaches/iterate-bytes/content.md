# Iterate as bytes

```go
// Package hamming is a small library for determining the hamming distance between two strands.
package hamming

import "errors"

// Distance accepts two strands and returns their hamming distance.
func Distance(strand1, strand2 string) (distance int, err error) {
	if len(strand1) != len(strand2) {
		return 0, errors.New("strands must be of equal length")
	}

	for i := 0; i < len(strand1); i++ {
		if strand1[i] != strand2[i] {
			distance++
		}
	}
	return distance, nil
}
```

The `Distance()` function is defined with [named return values][named-return-values] which are initialized with their [zero values][zero-values].

The [`len()`][len] function is used to compare the lengths of the strands in bytes.
Since at the time of writing all of the characters are [ASCII][ascii], getting the length of bytes is the same as getting the number of characters.
If the strands are not the same length, then `0` is returned along with the new error.

The ways to iterate characters are by Unicode runes, or by each letter being a string, or by each letter being a byte.
The runes are from [`range`][range] on a string, the strings from [`Split()`][split], and the bytes from indexing into the string.
Another way to iterate runes is to convert the string to a rune [`slice`][slice] and `range` on it.
The difference between ranging on a rune slice vs ranging on a string is that the index returned from a string is the position of the next rune in bytes,
not which rune it is.
For example, if the first unicode character is two bytes,
then the second unicode character index will be 2 when ranging on a string and 1 when ranging on a rune slice.
As of the time of this writing we can iterate bytes, since all of the characters are ASCII.

The `for` loop iterates from `0` until the length of the first strand.
Indexing into a string gets the byte at that position.
If the bytes are different at the same index for each strand, then the distance is incremented.
After the loop is finished, `distance` is returned along with a `nil` error.

[named-return-values]: https://yourbasic.org/golang/named-return-values-parameters/
[zero-values]: https://yourbasic.org/golang/default-zero-value/
[len]: https://pkg.go.dev/builtin#len
[ascii]: https://www.asciitable.com/
[range]: https://gobyexample.com/range
[split]: https://pkg.go.dev/strings#Split
[slice]: https://gobyexample.com/slices
