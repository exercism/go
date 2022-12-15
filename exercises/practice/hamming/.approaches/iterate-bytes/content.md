# Iterate as bytes

```go
// Package hamming is a small library for determining the hamming distance between two strands.
package hamming

import "errors"

// Distance accepts two strands and returns their hamming distance.
func Distance(strand1, strand2 string) (distance int, e error) {

	if len(strand1) != len(strand2) {
		return distance, errors.New("strands must be of equal length")
	}
	for i := 0; i < len(strand1); i++ {
		if strand1[i] != strand2[i] {
			distance++
		}
	}
	return distance, e
}
```

The `Distance()` function is defined with [named return values][named-return-values] which are initialized with their [zero values][zero-values].

The [`len()`][len] function is used to compare the lengths of the strands in bytes.
Since at the time of writing all of the characters are [ASCII][ascii], getting the length of bytes is the same as getting the number of characters.
If the strands are not the same length, then `distance` is returned (with its initial value of `0`) along with the new error.

The `for` loop iterates from `0` until the length of the first strand.
Indexing into a string gets the byte at that position.
If the bytes are different at the same index for each strand, then the distance is incremented.
After the loop is finished, `distance` is returned along with `e` (with its initial value of `nil`).

[named-return-values]: https://yourbasic.org/golang/named-return-values-parameters/
[zero-values]: https://yourbasic.org/golang/default-zero-value/
[len]: https://pkg.go.dev/builtin#len
[ascii]: https://www.asciitable.com/
