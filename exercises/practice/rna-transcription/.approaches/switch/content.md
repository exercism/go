# `switch` statement

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

This approach starts be defining a helper function which contains a [`switch`][switch] statement to map the DNA byte to an RNA byte.
The `ToRNA()` function begins by getting the [`len()`][len] of the input string in [bytes][byte].
A string can be seen as [bytes or runes][bytes-runes].
As of the time of this writing, all of the characters in the input `string` are single-byte [ASCII][ascii] characters,
so they can be treated as bytes.
A byte [slice][slice] is defined using [make][make], which is passed the length of the input `string` in bytes.

The ways to iterate characters are by Unicode runes, or by each letter being a string, or by each letter being a byte.
The runes are from [`range`][range] on a string, the strings from [`Split()`][split], and the bytes from indexing into the string.
Another way to iterate runes is to convert the string to a rune [`slice`][slice] and `range` on it.
The difference between ranging on a rune slice vs ranging on a string is that the index returned from a string is the position of the next rune in bytes,
not which rune it is.
For example, if the first unicode character is two bytes,
then the second unicode character index will be 2 when ranging on a string and 1 when ranging on a rune slice.
As of the time of this writing we can iterate bytes, since all of the characters are ASCII.

This iterates over the bytes of the input `string` and passes them as the argument to the helper function.
The helper function returns the matching RNA byte for the DNA byte, which is placed at the same index in the byte slice.
Although the values used in the `switch` are rune literals, they are treated as bytes, since their values are within the byte range,
and the type of the argument and return type are `byte`.

After the iteration of the input `string` finishes, the function returns the byte slice converted to a `string`.

[switch]: https://go.dev/tour/flowcontrol/9
[byte]: https://zetcode.com/golang/byte/
[len]: https://pkg.go.dev/builtin#len
[bytes-runes]: https://go.dev/blog/strings
[ascii]: https://www.asciitable.com/
[make]: https://go.dev/tour/moretypes/13
[range]: https://gobyexample.com/range
[split]: https://pkg.go.dev/strings#Split
[slice]: https://gobyexample.com/slices
