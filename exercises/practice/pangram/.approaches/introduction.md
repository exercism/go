# Introduction

There are at least two idomatic approaches to solving Pangram in Go.
You can use the `strings.ContainsRune()` function.
Or you can use a bit field to keep track of used letters.

## General guidance

The key to solving Pangram is determining if all of the letters in the alphabet are in the `string` being tested.
The occurrence of either the letter `a` or the letter `A` would count as the same letter.

## Approach: `strings.ContainsRune()`

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

For more information, check the [`strings.ContainsRune()` approach][approach-strings-containsrune].

## Approach: Bit field

```go
// Package pangram is a small library for checking if a phrase is a pangram
package pangram

const alphamask int = 0b11111111111111111111111111

// IsPangram checks if a phrase is a pangram
func IsPangram(phrase string) bool {
	phrasemask := 0
	length := len(phrase)
	for i := 0; i < length; i++ {
		letter := phrase[i]
		if letter > 96 && letter < 123 {
			phrasemask |= 1 << (letter - 97)
		} else if letter > 64 && letter < 91 {
			phrasemask |= 1 << (letter - 65)
		}
	}
	return alphamask&phrasemask == alphamask
}
```

For more information, check the [Bit field approach][approach-bitfield].

## Which approach to use?

The bit field approach benchmarked the fastest.
To compare performance of the approaches, check the [Performance article][article-performance].

The `strings.ContainsRune()` runes approach is much more readable so if the code is not extremely performance critical, then opting for a readable, easy-to-change solution is recommended.

[approach-strings-containsrune]: https://exercism.org/tracks/go/exercises/pangram/approaches/strings-containsrune
[approach-bitfield]: https://exercism.org/tracks/go/exercises/pangram/approaches/bitfield
[article-performance]: https://exercism.org/tracks/go/exercises/pangram/articles/performance
