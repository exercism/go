# `strings.ContainsRune()`

```go
// Package pangram is a small library for checking if a phrase is a pangram
package pangram

import "strings"

// IsPangram checks if a phrase is a pangram
func IsPangram(s string) bool {
	lookup := strings.ToLower(s)
	for chr := 'a'; chr <= 'z'; chr++ {
		if !strings.ContainsRune(lookup, chr) {
			return false
		}
	}
	return true
}
```

This begins by lowercasing the input by using the [`strings`][strings] [`ToLower()`][tolower] function.
Since, for example, `A` is considered the same letter as `a`, lowercasing `A` makes it easier to check if `a` is in the input.

The [runes][rune] of `a` to `z` are iterated in the `for` loop.
The [strings.ContainsRune()][strings-containsrune] function is used to see if the letter is contained in the input `string`.
If not, then the function immediately returns `false`.

If the loop completes without returning `false`, then all of the a-z letters are contained in the input `string`,
and the function returns `true`.

[strings]: https://pkg.go.dev/strings
[tolower]: https://pkg.go.dev/strings#ToLower
[rune]: https://pkg.go.dev/builtin#rune
[strings-containsrune]: https://pkg.go.dev/strings#ContainsRune
