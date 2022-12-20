# Introduction

There are various idiomatic approaches to solve Scrabble Score.
One approach can use a [`map`][map] to look up the scores.
Or a [switch][switch] can be used to look up the scores.

## General guidance

It can be helpful to understand the difference between [strings of one letter and runes][strings-bytes-runes].

## Approach: `map`

```go
// Package scrabble scores a scrabble word.
package scrabble

import "unicode"

var lookup = map[rune]int{
	'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1, 'l': 1, 'n': 1, 'r': 1, 's': 1, 't': 1,
	'd': 2, 'g': 2,
	'b': 3, 'c': 3, 'm': 3, 'p': 3,
	'f': 4, 'h': 4, 'v': 4, 'w': 4, 'y': 4,
	'k': 5,
	'j': 8, 'x': 8,
	'q': 10, 'z': 10,
}

// Score takes a word and returns its scrabble score.
func Score(word string) (score int) {
	for _, ltr := range word {
		score += lookup[unicode.ToLower(ltr)]
	}
	return score
}
```

For more information, check the [`map` approach][approach-map].

## Approach: `switch` on runes

```go
// Package scrabble uses switch to score a scrabble word.
package scrabble

import "unicode"

// Score takes a word and returns its scrabble score.
func Score(word string) (score int) {

	for _, letter := range word {

		switch unicode.ToLower(letter) {
		case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
			score += 1
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		}
	}
	return score
}
```

For more information, check the [`switch` on runes approach][approach-switch-on-runes].

## Other approaches

Besides the aforementioned, idiomatic approaches, you could also approach the exercise as follows:

### Other approach: `switch` on strings

The input can be split into each character being a separate string.
A  `switch` can be used to look up the score for each single-letter string.

For more information, check the [`switch` on strings approach][approach-switch-on-strings].

## Which approach to use?

The `switch` on runes approach benchmarked as the fastest approach.

For more information, check the [Performance article][article-performance].

[map]: https://gobyexample.com/maps
[switch]: https://gobyexample.com/switch
[strings-bytes-runes]: https://go.dev/blog/strings
[approach-map]: https://exercism.org/tracks/go/exercises/scrabble-score/approaches/map
[approach-switch-on-runes]: https://exercism.org/tracks/go/exercises/scrabble-score/approaches/switch-on-runes
[approach-switch-on-strings]: https://exercism.org/tracks/go/exercises/scrabble-score/approaches/switch-on-strings
[article-performance]: https://exercism.org/tracks/go/exercises/scrabble-score/articles/performance
