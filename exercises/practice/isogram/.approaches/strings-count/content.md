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
- [`unicode.IsLetter()`][isletter] is used to ignore characters in the `string` that are not Unicode letters.
- [`strings.Count()`][count] is used to get the number of occurrences of the letter [`rune`][rune] converted to a [`string`][string].
- If the letter occurs more than once, then the function returns false.

If the loop finishes without returning `false`, then no letters are repeated, and the function returns `true`.

[tolower]: https://pkg.go.dev/strings#ToLower
[isletter]: https://pkg.go.dev/unicode#IsLetter
[rune]: https://pkg.go.dev/builtin#rune
[string]: https://pkg.go.dev/builtin#string
[count]: https://pkg.go.dev/strings#Count
