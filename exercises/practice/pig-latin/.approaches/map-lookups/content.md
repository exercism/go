# map lookups

```go
// Package piglatin is a small library for translating a sentence into Pig Latin.
package piglatin

import (
	"strings"
)

var vowels = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
var specials = map[string]bool{"xr": true, "yt": true}
var vowels_y = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true, 'y': true}

// Sentence translates a sentence into Pig Latin.
func Sentence(phrase string) string {
	piggyfied := strings.Builder{}

	for _, word := range strings.Split(phrase, " ") {
		if piggyfied.Len() != 0 {
			piggyfied.WriteByte(' ')
		}

		if vowels[word[0]] || specials[word[:2]] {
			piggyfied.WriteString(word + "ay")
			continue
		}

		for pos := 1; pos < len(word); pos++ {
			letter := word[pos]
			if vowels_y[letter] {
				if letter == 'u' && word[pos-1] == 'q' {
					pos++
				}
				piggyfied.WriteString(word[pos:] + word[:pos] + "ay")
				break
			}
		}
	}
	return piggyfied.String()
}
```

The approach imports the `strings` package so it can use [`strings.Builder`][builder] as an efficient way to build the output string.

This approach starts be defining several [map][map]s for looking up vowels and special letter combinations.
The maps have boolean values so that a key can be looked up simply as 

```go
if vowels[word[0]]
```

instead of 

```
if _, ok := vowels[word[0]]; ok {
```

Doing so also allows having multiple checks on the same line,
since there can not be multiple [`if` with short statements][if-with-short-statement] on the same line.

After instantiating a `strings.Builder` the sentence is [`Split`][split] with a space separator to iterate the words.
Unless it is the first word, a space is written to the `Builder` to separate the word from the previous word.

The word is checked for starting with a vowel or if the beginning [slice][slice] of two characters is in the map of "specials".
If so, the word is written to the `Builder` along with `ay` at the end.
Then the iteration is [continue][continue]d to the next word.
If not, each byte in the word is iterated from index 1 up to the length of the word, until finding a vowel (that now includes `y`).

The ways to iterate characters are by Unicode runes, or by each letter being a string, or by each letter being a byte.
The runes are from [`range`][range] on a string, the strings from [`Split()`][split], and the bytes from indexing into the string.
Another way to iterate runes is to convert the string to a rune [`slice`][slice] and `range` on it.
The difference between ranging on a rune slice vs ranging on a string is that the index returned from a string is the position of the next rune in bytes,
not which rune it is.
For example, if the first unicode character is two bytes,
then the second unicode character index will be 2 when ranging on a string and 1 when ranging on a rune slice.
As of the time of this writing we can iterate bytes, since all of the characters are single-byte [ASCII][ascii] characters.

If the found vowel is `u`, the previous consonant is checked for being `q` and the position is adjusted accordingly.
A slice is then taken from the position until the end of the word,
plus a slice taken from the beginning of the word and stopped before the position, plus `ay`, and added to the `Builder`.

After the iteraton of words is finished, the `Builder`'s [`String`][string] method is called to return the accumulated string from the function.

[strings]: https://pkg.go.dev/strings
[builder]: https://pkg.go.dev/strings#Builder
[map]: https://gobyexample.com/maps
[if-with-short-statement]: https://go.dev/tour/flowcontrol/6
[split]: https://pkg.go.dev/strings#Split
[slice]: https://go.dev/tour/moretypes/7
[continue]: https://www.tutorialspoint.com/go/go_continue_statement.htm
[ascii]: https://www.asciitable.com/
[range]: https://gobyexample.com/range
[string]: https://pkg.go.dev/strings#Builder.String
