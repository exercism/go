# Introduction

There are various idiomatic ways to solve Isogram.
You can use a [`map`][map] to keep track of used letters.
You can use the [`strings`][strings] function [`Count()`][strings-count] to calculate used letters.
Or you could use a bit field to keep track of used letters.

## General guidance

The key to solving Isogram is to determine if any of the letters in the `string` being checked are repeated.
A repeated letter means the `string` is not an Isogram.
The occurrence of the letter `a` and the letter `A` count as a repeated letter, so `Alpha` would not be an isogram.

## Approach: `map`

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
		_, exists := lookup[ltr]
		if exists {
			return false
		}
		lookup[ltr] = blank{}
	}
	return true
}
```

For more information, check the [`map` approach][approach-map].

## Approach: `strings.Count()`

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

For more information, check the [`strings.Count()` approach][approach-strings-count].

## Approach: Bit field


```go
// Package isogram is a small library for analyzing if a phrase is an isogram.
package isogram

// IsIsogram accepts a phrase and returns if it is an isogram.
func IsIsogram(phrase string) bool {
	theEnd := len(phrase)
	var letterFlags uint32 = 0

	for i := 0; i < theEnd; i++ {
		letter := phrase[i]
		if letter > 96 && letter < 123 {
			if (letterFlags & (1 << (letter - 'a'))) != 0 {
				return false
			} else {
				letterFlags |= (1 << (letter - 'a'))
			}
		} else if letter > 64 && letter < 91 {
			if (letterFlags & (1 << (letter - 'A'))) != 0 {
				return false
			} else {
				letterFlags |= (1 << (letter - 'A'))
			}
		}
	}
	return true
}
```

For more information, check the [Bit field approach][approach-bitfield].

## Which approach to use?

The fastest is the `Bit field` approach, but it depends on all of the letters being [ASCII][ascii].
The fastest approach that supports Unicode is the `strings.Count()` approach.

To compare performance of the approaches, check the [Performance article][article-performance].

[map]: https://gobyexample.com/maps
[strings]: https://pkg.go.dev/strings
[strings-count]: https://pkg.go.dev/strings#Count
[approach-map]: https://exercism.org/tracks/go/exercises/isogram/approaches/map
[approach-strings-count]: https://exercism.org/tracks/go/exercises/isogram/approaches/strings-count
[approach-bitfield]: https://exercism.org/tracks/go/exercises/isogram/approaches/bitfield
[article-performance]: https://exercism.org/tracks/go/exercises/isogram/articles/performance
[ascii]: https://www.asciitable.com/
