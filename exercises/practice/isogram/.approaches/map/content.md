# `map`

```go
// Package isogram is a small library for analyzing if a phrase is an isogram.
package isogram

import (
	"unicode"
)

type blank struct{}

// IsIsogram accepts a phrase and returns if it is an isogram.
func IsIsogram(phrase string) bool {
	lookup := make(map[rune]blank)
	for _, chr := range phrase {
		if !unicode.IsLetter(chr) {
			continue
		}
		ltr := unicode.ToLower(chr)
		if _, exists := lookup[ltr]; exists {
			return false
		}
		lookup[ltr] = blank{}
	}
	return true
}
```

This approach starts by defining the type `blank` as an empty struct.
The lookup [`map`][map] is defined with a [`rune`][rune] as the key and `blank` (i.e. an empty struct) as the value.

~~~~exercism/note
By assigning values of an empty struct to the `map`, it essentially makes the keys what is known as a "set" in other languages.
A set is a collection of unique values.
More info on implementing a set in Go can be found [here](https://yourbasic.org/golang/implement-set/).
~~~~

The ways to iterate characters are by Unicode runes, or by each letter being a string, or by each letter being a byte.
The runes are from [`range`][range] on a string, the strings from [`Split()`][split], and the bytes from indexing into the string.
Another way to iterate runes is to convert the string to a rune [`slice`][slice] and `range` on it.
The difference between ranging on a rune slice vs ranging on a string is that the index returned from a string is the position of the next rune in bytes,
not which rune it is.
For example, if the first unicode character is two bytes, then the second unicode character index will be 2 when ranging on a string and 1 when ranging on a rune slice.
In this approach we don't care about the index, so ranging on a `string` is fine.

- [`unicode.IsLetter()`][isletter] is used to ignore the character if it is not a Unicode letter.
- [`unicode.ToLower()`][tolower] is used to ensure the letter is lowercase.
The occurrence of the letter `a` and the letter `A` count as a repeated letter, so `Alpha` would not be an isogram.
Lowercasing `A` to `a` makes it easier to check for the repeated `a`.

- An [`if` with a short statement][if-assignment] is used to both assign the `exists` variable and to test if it's `true`.
If it is `true`, then the letter has been used before, and the function returns `false`.
If it is not `true`, the control falls through to the next statement, which creates the letter key and assigns it the empty struct value.

If the loop finishes without returning `false`, then no letters are repeated, and the function returns `true`.

[map]: https://gobyexample.com/maps
[rune]: https://pkg.go.dev/builtin#rune
[range]: https://gobyexample.com/range
[split]: https://pkg.go.dev/strings#Split
[slice]: https://gobyexample.com/slices
[isletter]: https://pkg.go.dev/unicode#IsLetter
[tolower]: https://pkg.go.dev/unicode#ToLower
[if-assignment]: https://go.dev/tour/flowcontrol/6
