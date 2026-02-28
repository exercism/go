# `strings.Map()` function

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

This approach passes the input `string` into the [`strings`][strings] [`Map()`][strings-map] function.
The function uses a [`switch`][switch] statement to translate the DNA character into an RNA character.

[strings]: https://pkg.go.dev/strings
[strings-map]: https://pkg.go.dev/strings#Map
[switch]: https://go.dev/tour/flowcontrol/9
