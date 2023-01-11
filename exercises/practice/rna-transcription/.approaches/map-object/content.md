# `map` object

```go
// Package strand is a small library for returning the RNA complement of a DNA strand.
package strand

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

This approach starts be defining a [`map`][map] to translate the DNA [rune] to an RNA [byte].
It does this outside of the `ToRNA()` function, so it is not created every time `ToRNA()` is called.
A byte [slice][slice] is defined using [make][make], which is passed the [`len()`][len] of the input `string` in bytes.
A string can be seen as [bytes or runes][bytes-runes].
As of the time of this writing, all of the characters in the input `string` are single-byte [ASCII][ascii] characters,
so they can be treated as bytes.

The ways to iterate characters are by Unicode runes, or by each letter being a string, or by each letter being a byte.
The runes are from [`range`][range] on a string, the strings from [`Split()`][split], and the bytes from indexing into the string.
Another way to iterate runes is to convert the string to a rune [`slice`][slice] and `range` on it.
The difference between ranging on a rune slice vs ranging on a string is that the index returned from a string is the position of the next rune in bytes,
not which rune it is.
For example, if the first unicode character is two bytes,
then the second unicode character index will be 2 when ranging on a string and 1 when ranging on a rune slice.

This iterates over the runes of the input `string` and passes them as the key to the `map`.
For this exercise we are not concerned with validating the letter and handling an error for an illegal character.
The `map` returns the matching RNA byte value for the DNA rune, which is placed at the same index in the byte slice.

```exercism/caution
This iterates over the runes of the input `string` and places them at the same index in the byte slice.
That works in this case, because the input is all single-byte ASCII characters.
But, if the input included any multi-byte Unicode characters, the indexes would not match.
```

After the iteration of the input `string` finishes, the function returns the byte slice converted to a `string`.

[map]: https://gobyexample.com/maps
[rune]: https://golangdocs.com/rune-in-golang
[byte]: https://zetcode.com/golang/byte/
[make]: https://go.dev/tour/moretypes/13
[len]: https://pkg.go.dev/builtin#len
[bytes-runes]: https://go.dev/blog/strings
[ascii]: https://www.asciitable.com/
[range]: https://gobyexample.com/range
[split]: https://pkg.go.dev/strings#Split
[slice]: https://gobyexample.com/slices
