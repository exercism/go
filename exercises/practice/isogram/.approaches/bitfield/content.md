# Bit field

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

This solution uses the [ASCII][ascii] value of the letter to set the corresponding bit position.

The ways to iterate characters are by Unicode runes, or by each letter being a string, or by each letter being a byte.
The runes are from [`range`][range] on a string, the strings from [`Split()`][split], and the bytes from indexing into the string.
Another way to iterate runes is to convert the string to a rune [`slice`][slice] and `range` on it.
The difference between ranging on a rune slice vs ranging on a string is that the index returned from a string is the position of the next rune in bytes,
not which rune it is.
For example, if the first unicode character is two bytes, then the second unicode character index will be 2 when ranging on a string and 1 when ranging on a rune slice.
As of the time of this writing we can iterate bytes, since all of the characters are ASCII.

The `string` is looped through its characters, looking for a character being `a` through `z` or `A` through `Z`.
- The ASCII value for `a` is `97`, and for `z` is `122`.
- The ASCII value for `A` is `65`, and for `Z` is `90`.

- If the lower-cased letter is subtracted by `a`, then `a` will result in `0`, because `97` minus `97`  equals `0`.
`z` would result in `25`, because `122` minus `97` equals `25`.
So `a` would have `1` [shifted left][shift-left] 0 places (so not shifted at all) and `z` would have `1` shifted left 25 places.
- If the upper-cased letter is subtracted by `A`, then `A` will result in `0`, because `65` minus `65`  equals `0`.
`Z` would result in `25`, because `90` minus `65` equals `25`.
So `A` would have `1` [shifted left][shift-left] 0 places (so not shifted at all) and `Z` would have `1` shifted left 25 places.

In that way, both a lower-cased `z` and an upper-cased `Z` can share the same position in the bit field.

So, for an unsigned thirty-two bit integer, if the values for `a` and `Z` were both set, the bits would look like

```
      zyxwvutsrqponmlkjihgfedcba
00000010000000000000000000000001
```

We can use the [bitwise AND operator][and] to check if a bit has already been set.
If it has been set, we know the letter is duplicated and we can immediately return `false`.
If it has not been set, we can use the [bitwise OR operator][or] to set the bit.
If the loop completes without finding a duplicate letter (and returning `false`), the function returns `true`.

This approach is fast, but it only works with the `ASCII` alphabet.

[ascii]: https://www.asciitable.com/
[range]: https://gobyexample.com/range
[split]: https://pkg.go.dev/strings#Split
[slice]: https://gobyexample.com/slices
[shift-left]: https://www.golangprograms.com/bitwise-operators-in-go-programming-language.html
[or]: https://yourbasic.org/golang/bitwise-operator-cheat-sheet/
[and]: https://yourbasic.org/golang/bitwise-operator-cheat-sheet/
