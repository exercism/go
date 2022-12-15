# Introduction

There are at least two idomatic approaches to solving Hamming in Go.
You can iterate the characters as [bytes][bytes].
Or you can iterate the characters as [runes][runes].

## General guidance

One important aspect to solving Hamming is to verify the two strands have the same number of characters.
Another aspect to be aware of is to iterate and compare the two input string characters consistently as either byte to byte or rune to rune.

## Approach: Iterate as bytes

```go
// Package hamming is a small library for determining the hamming distance between two strands.
package hamming

import "errors"

// Distance accepts two strands and returns their hamming distance.
func Distance(strand1, strand2 string) (distance int, e error) {

	if len(strand1) != len(strand2) {
		return distance, errors.New("strands must be of equal length")
	}
	for i := 0; i < len(strand1); i++ {
		if strand1[i] != strand2[i] {
			distance++
		}
	}
	return distance, e
}
```

For more information, check the [iterate as bytes approach][approach-iterate-bytes].

## Approach: Iterate as runes

```go
// Package hamming is a small library for determining the hamming distance between two strands.
package hamming

import (
	"errors"
	"unicode/utf8"
)

// Distance accepts two strands and returns their hamming distance.
func Distance(strand1 string, strand2 string) (distance int, e error) {
	if utf8.RuneCountInString(strand1) != utf8.RuneCountInString(strand2) {
		return distance, errors.New("strands must be of equal length")
	}

	strand1Runes := []rune(strand1)
	strand2Runes := []rune(strand2)
	for idx, rooney := range strand1Runes {
		if rooney != strand2Runes[idx] {
			distance++
		}
	}
	return distance, e
}
```

For more information, check the [iterate as runes approach][approach-iterate-runes].

## Which approach to use?

To iterate as bytes benchmarked the fastest.
As of the time of this writing we can iterate bytes, since all of the characters are [ASCII][ascii].

To compare performance of the approaches, check the [Performance article][article-performance].

[bytes]: https://pkg.go.dev/bytes
[runes]: https://golangdocs.com/rune-in-golang
[approach-iterate-bytes]: https://exercism.org/tracks/go/exercises/hamming/approaches/iterate-bytes
[approach-iterate-runes]: https://exercism.org/tracks/go/exercises/hamming/approaches/iterate-runes
[article-performance]: https://exercism.org/tracks/go/exercises/hamming/articles/performance
[ascii]: https://www.asciitable.com/
