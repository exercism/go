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

This approach starts be defining several [map][]map]s for looking up vowels and special letter combinations.
The maps have boolean values so that a key can be looked up simply as 

```go
if vowels[word[0]]
```

instead of 

```
if _, ok := vowels[word[0]]; ok {
```

Doing so also allows having multiple checks on the same line.
There can not be mulitple [if with short statements][if-with-short-statement] on the same line.

TODO

[strings]: https://pkg.go.dev/strings
[builder]: https://pkg.go.dev/strings#Builder
[map]: https://gobyexample.com/maps
[if-with-short-statement]: https://go.dev/tour/flowcontrol/6
