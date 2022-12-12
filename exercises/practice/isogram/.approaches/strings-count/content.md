# `strings.Count()`

```go
// Package isogram is a small library for analyzing if a phrase is an isogram.
package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram accepts a phrase and returns if it is an isogram.
func IsIsogram(phrase string) bool {
	phrase = strings.ToLower(phrase)
	for _, chr := range phrase {
		if !unicode.IsLetter(chr) {
			continue
		}
		if strings.Count(phrase, string(chr)) > 1 {
			return false
		}
	}
	return true
}
```


The steps for this solution are

- [`strings.ToLower()`][tolower] produces a new `string` with the characters of the passed-in `string` changed to lower case.
The occurrence of the letter `a` and the letter `A` count as a repeated letter, so `Alpha` would not be an isogram.
Lowercasing `Alpha` to `alpha` makes it easier to check for the repeated `a`.

The ways to iterate characters are by Unicode runes, or by each letter being a string, or by each letter being a byte.
The runes are from [`range`][range] on a string, the strings from [`Split()`][split], and the bytes from indexing into the string.
Another way to iterate runes is to convert the string to a rune [`slice`][slice] and `range` on it.
The difference between ranging on a rune slice vs ranging on a string is that the index returned from a string is the position of the next rune in bytes,
not which rune it is.
For example, if the first unicode character is two bytes, then the second unicode character index will be 2 when ranging on a string and 1 when ranging on a rune slice.
In this approach we don't care about the index, so ranging on a `string` is fine.

- [`unicode.IsLetter()`][isletter] is used to ignore the character if it is not a Unicode letter.
- The letter [`rune`][rune] is converted to a [`string`][string] to be passed to `strings.Count()`.
[`strings.Count()`][count] is used to get the number of occurrences of the letter in the entire `string`.
- If the letter occurs more than once, then the function returns false.

If the loop finishes without returning `false`, then no letters are repeated, and the function returns `true`.

[tolower]: https://pkg.go.dev/strings#ToLower
[isletter]: https://pkg.go.dev/unicode#IsLetter
[rune]: https://pkg.go.dev/builtin#rune
[range]: https://gobyexample.com/range
[split]: https://pkg.go.dev/strings#Split
[slice]: https://gobyexample.com/slices
[string]: https://pkg.go.dev/builtin#string
[count]: https://pkg.go.dev/strings#Count
