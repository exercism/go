# Introduction

There are various ways to solve Pig Latin.
One way is to use [regular expressions][regular-expressions] (also known as [regex][regex]) for processing the input.
Solutions using regex require familiarity with regex patterns, which are like another language.
Another way is to use a series of conditional statements to test which of several rules the input matches.
Another approach is to use [map][map]s for look-up and use either an index into or a [slice][slice] of the input as a key for the map.

## General guidance

At the time of writing only four rules need to be handled, but if they have similar output, they don't need to be handled completely separately.

## Approach: map lookups

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

For more information, check the [map lookups approach][approach-map-lookups].

## Approach: map lookups with generics

```go
// Package piglatin is a small library for translating a sentence into Pig Latin.
package piglatin

import (
	"strings"
)

type void struct{}

var blank void
var vowels = map[byte]void{'a': blank, 'e': blank, 'i': blank, 'o': blank, 'u': blank}
var specials = map[string]void{"xr": blank, "yt": blank}
var vowels_y = map[byte]void{'a': blank, 'e': blank, 'i': blank, 'o': blank, 'u': blank, 'y': blank}

type container interface {
	byte | string
}

func contains[T container](values map[T]void, value T) bool {
	_, ok := values[value]
	return ok
}

// Sentence translates a sentence into Pig Latin.
func Sentence(phrase string) string {
	piggyfied := strings.Builder{}

	for _, word := range strings.Split(phrase, " ") {
		if piggyfied.Len() != 0 {
			piggyfied.WriteByte(' ')
		}

		if contains(vowels, word[0]) || contains(specials, word[:2]) {
			piggyfied.WriteString(word + "ay")
			continue
		}

		for pos := 1; pos < len(word); pos++ {
			letter := word[pos]
			if contains(vowels_y, letter) {
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

For more information, check the [map lookups with generics approach][approach-map-lookups-with-generics].

## Which approach to use?

The map lookups approach has less boilerplate, so it may be considered more readable.
The fastest benchmarks were from map lookups with generics, but not consistently so.
Benchmarks runs for each approach varied by as much as 100 nanoseconds to 200 nanoseconds or so, but the fastest benchmarks from several runs were for map lookups with generics.
To compare performance of the approaches, check the [Performance article][article-performance].

[regular-expressions]: https://gobyexample.com/regular-expressions
[regex]: https://pkg.go.dev/regexp
[map]: https://gobyexample.com/maps
[slice]: https://gobyexample.com/slices
[approach-map-lookups]: https://exercism.org/tracks/go/exercises/pig-latin/approaches/map-lookups
[approach-map-lookups-with-generics]: https://exercism.org/tracks/go/exercises/pig-latin/approaches/map-lookups-with-generics
[article-performance]: https://exercism.org/tracks/go/exercises/pig-latin/articles/performance
