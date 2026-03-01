# Introduction

There are at least three idomatic approaches to solving RNA Transcription in Go.
You can use the `strings.Map()` function.
Or you can use the `map` object to look up the letters.
Or a `switch` statement can be used to translate the letters.

## General guidance

The key to solving RNA Transcription is to match the RNA value to the DNA value.

## Approach: `strings.Map()` function

```go
// Package strand is a small library for returning the RNA complement of a DNA strand.
package rnatranscription

import "strings"

// toRNA returns the RNA complement of the input DNA strand.
func ToRNA(dna string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case 'G':
			return 'C'
		case 'C':
			return 'G'
		case 'T':
			return 'A'
		case 'A':
			return 'U'
		}
		return r
	}, dna)
}
```

For more information, check the [`strings.Map()` approach][approach-strings-map].

## Approach: `map` object

```go
// Package strand is a small library for returning the RNA complement of a DNA strand.
package rnatranscription

var translate = map[rune]byte{
	'G': 'C',
	'C': 'G',
	'T': 'A',
	'A': 'U',
}

// ToRNA returns the RNA complement of the input DNA strand.
func ToRNA(dna string) string {
	rna := make([]byte, len(dna))
	for i, c := range dna {
		rna[i] = translate[c]
	}
	return string(rna)
}
```

For more information, check the [`map` object approach][approach-map-object].

## Approach: `switch` statement

```go
// Package strand is a small library for returning the RNA complement of a DNA strand.
package rnatranscription

func nucComp(nuc byte) byte {
	switch nuc {
	case 'G':
		return 'C'
	case 'C':
		return 'G'
	case 'T':
		return 'A'
	case 'A':
		return 'U'
	default:
		panic(string(nuc) + " is an invalid nucleotide")
	}
}

// toRNA returns the RNA complement of the input DNA strand.
func ToRNA(dna string) string {
	length := len(dna)
	var output = make([]byte, length)
	for i := 0; i < length; i++ {
		output[i] = (nucComp(dna[i]))
	}
	return string(output)
}
```

For more information, check the [`switch` statement approach][approach-switch-statement]

## Which approach to use?

The `switch` statement approach benchmarked the fastest.
To compare performance of the approaches, check the [Performance article][article-performance].

[approach-strings-map]: https://exercism.org/tracks/go/exercises/rna-transcription/approaches/map-function
[approach-map-object]: https://exercism.org/tracks/go/exercises/rna-transcription/approaches/map-object
[approach-switch-statement]: https://exercism.org/tracks/go/exercises/rna-transcription/approaches/switch
[article-performance]: https://exercism.org/tracks/go/exercises/rna-transcription/articles/performance
