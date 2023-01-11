# `switch` on strings

```go
// Package scrabble uses switch to score a scrabble word.
package scrabble

import "strings"

// Score takes a word and returns its scrabble score.
func Score(word string) (score int) {
	for _, letter := range strings.Split(strings.ToLower(word), "") {

		switch letter {
		case "a", "e", "i", "o", "u", "l", "n", "r", "s", "t":
			score += 1
		case "d", "g":
			score += 2
		case "b", "c", "m", "p":
			score += 3
		case "f", "h", "v", "w", "y":
			score += 4
		case "k":
			score += 5
		case "j", "x":
			score += 8
		case "q", "z":
			score += 10
		}
	}
	return score
}
```

The `Score()` function is defined with a [named return value][named-return-value] which is initialized with its [zero value][zero-value].

The ways to iterate characters are by Unicode runes, or by each letter being a string, or by each letter being a byte.
The runes are from [`range`][range] on a string, the strings from [`Split()`][split], and the bytes from indexing into the string.
Another way to iterate runes is to convert the string to a rune [`slice`][slice] and `range` on it.
The difference between ranging on a rune slice vs ranging on a string is that the index returned from a string is the position of the next rune in bytes,
not which rune it is.
For example, if the first unicode character is two bytes, then the second unicode character index will be 2 when ranging on a string and 1 when ranging on a rune slice.
As of the time of this writing we are only iterating ASCII characters, so we don't care about the index, and ranging on the `Split()` input string is fine.

For this exercise we are not concerned with validating the letter and handling an error for an illegal character.
The `strings.ToLower()` function is used to ensure all of the letters are lowercase.
Lowercase letters are used instead of uppercase because most of the input will be lowercase.
It makes a slight performance difference, because of how [`strings.ToLower()`][tolower] is implemented, as seen in the strings.go file:

```go
// ToLower returns s with all Unicode letters mapped to their lower case.
func ToLower(s string) string {
	isASCII, hasUpper := true, false
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= utf8.RuneSelf {
			isASCII = false
			break
		}
		hasUpper = hasUpper || ('A' <= c && c <= 'Z')
	}

	if isASCII { // optimize for ASCII-only strings.
		if !hasUpper {
			return s
		}
		var b Builder
		b.Grow(len(s))
		for i := 0; i < len(s); i++ {
			c := s[i]
			if 'A' <= c && c <= 'Z' {
				c += 'a' - 'A'
			}
			b.WriteByte(c)
		}
		return b.String()
	}
	return Map(unicode.ToLower, s)
}
```

We see that when the rune is an [ASCII][ascii] character it is only lowercased if it is not already lowercase.
So, when most of the input is already lowercase, we can save a bit of time by only looking up the lowercase letters.
We could use both lowercase and uppercase letters in the `switch`, but that's a lot of typing!
Another consideration is that if the score of a letter were ever to change,
if we look up both lowercase and uppercase letters then we would have to change the score in two places.
It's preferable to have only one place to change something if something gets modified.
Some people may find uppercased letters to be more readable in this context, and they may take the slight performance hit to use those instead.

Each letter is used as the value to look up the score in the [`switch`][switch] statement,
which is added to the output variable.

When the iteration is finished, the output variable of the totaled score is returned from the function.

[tolower]: https://pkg.go.dev/strings#ToLower
[rune]: https://pkg.go.dev/builtin#rune
[ascii]: https://www.asciitable.com/
[named-return-value]: https://yourbasic.org/golang/named-return-values-parameters/
[zero-value]: https://yourbasic.org/golang/default-zero-value/
[range]: https://gobyexample.com/range
[split]: https://pkg.go.dev/strings#Split
[slice]: https://gobyexample.com/slices
[switch]: https://go.dev/tour/flowcontrol/9
