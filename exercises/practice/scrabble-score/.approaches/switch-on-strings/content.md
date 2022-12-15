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

in strings.go:

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
