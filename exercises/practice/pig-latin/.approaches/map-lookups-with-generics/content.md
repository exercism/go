# map lookups with generics

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

The approach imports the `strings` package so it can use [`strings.Builder`][builder] as an efficient way to build the output string.

This approach starts be defining a void [type][type] that will be used in the package.

Several [map][map]s for then defined for looking up vowels and special letter combinations.
The maps have an empty struct as the type for their values.

~~~~exercism/note
By assigning values of an empty struct to the `map`, it essentially makes the keys what is known as a "set" in other languages.
A set is a collection of unique values.
More info on implementing a set in Go can be found [here](https://yourbasic.org/golang/implement-set/).
~~~~

An [interface type][interface-type] is then defined which will be used to constain which types are to be used for the [generics][generics]:

> Generics are a way of writing code that is independent of the specific types being used. Functions and types may now be written to use any of a set of types.

Since there can not be multiple [`if` with short statements][if-with-short-statement] on the same line,
the `contains()` function is defined which takes a generic type parameter of `[T container]`.
`T` is the name of the parameter.
The name does not have to be `T`, but it is a common convention in Go (and other languages) to start with `T`,
and if a second name is needed, to use `U`, and so on.
The type of `T` is constrained to be one of the types defined in the `container` interface.
The arguments to the function are a map of the generic type `T` with void values, and a value of type `T`.
So, `contains()` can take arguments of either `map[byte]void, byte` or `map[string]void, string`.
The `contains()` function is defined to return a boolean, which is whether the value is a key of the map.

After instantiating a `strings.Builder` the sentence is [`Split`][split] with a space separator to iterate the words.
Unless it is the first word, a space is written to the `Builder` to separate the word from the previous word.

The `contains()` function is used to check if the word starts with a vowel or if the beginning [slice][slice] of two characters is in the map of "specials".
If so, the word is written to the `Builder` along with `ay` at the end.
Then the iteration is [continue][continue]d to the next word.
If not, each byte in the word is iterated from index 1 up to the length of the word, until finding a vowel (that now includes `y`).

The ways to iterate characters are by Unicode runes, or by each letter being a string, or by each letter being a byte.
The runes are from [`range`][range] on a string, the strings from [`Split()`][split], and the bytes from indexing into the string.
Another way to iterate runes is to convert the string to a rune [`slice`][slice] and `range` on it.
The difference between ranging on a rune slice vs ranging on a string is that the index returned from a string is the position of the next rune in bytes,
not which rune it is.
For example, if the first unicode character is two bytes,
then the second unicode character index will be 2 when ranging on a string and 1 when ranging on a rune slice.
As of the time of this writing we can iterate bytes, since all of the characters are single-byte [ASCII][ascii] characters.

If the found vowel is `u`, the previous consonant is checked for being `q` and the position is adjusted accordingly.
A slice is then taken from the position until the end of the word,
plus a slice taken from the beginning of the word and stopped before the position, plus `ay`, and added to the `Builder`.

After the iteraton of words is finished, the `Builder`'s [`String`][string] method is called to return the accumulated string from the function.

[strings]: https://pkg.go.dev/strings
[builder]: https://pkg.go.dev/strings#Builder
[types]: https://go.dev/ref/spec#Types
[map]: https://gobyexample.com/maps
[interface-type]: https://go.dev/ref/spec#Interface_types
[generics]: https://go.dev/blog/intro-generics
[if-with-short-statement]: https://go.dev/tour/flowcontrol/6
[split]: https://pkg.go.dev/strings#Split
[slice]: https://go.dev/tour/moretypes/7
[continue]: https://www.tutorialspoint.com/go/go_continue_statement.htm
[ascii]: https://www.asciitable.com/
[range]: https://gobyexample.com/range
[string]: https://pkg.go.dev/strings#Builder.String
