# Iterate as bytes

```go
// Package isbn is a small library for validating input as an isbn.
package isbn

func getDigit(chr byte) (digit int, ok bool) {
	if chr < '0' || chr > '9' {
		return digit, ok
	}
	return int(chr - '0'), true
}

// IsValidISBN calculates the validity of the input as an ISBN.
func IsValidISBN(input string) bool {

	length := len(input)
	pos := 10
	sum := 0

	for i := 0; i < length; i++ {
		chr := input[i]
		if digit, ok := getDigit(chr); ok {
			sum += (digit * pos)
			pos--
		} else if pos == 1 && chr == 'X' {
			sum += 10
			pos--
		} else if chr != '-' {
			return false
		}
	}
	return pos == 0 && sum%11 == 0
}
```

This approach starts be defining a helper function to return an `int` for a [`byte`][bytes] if the `byte` is the [ASCII][ascii]  value for a number.
The helper function is defined with [named return values][named-return-values] which are initialized with their [zero values][zero-values].
If the `byte` is not an ASCII number it immediately returns `0, false` to indicate that the `byte` is not a digit.
Otherwise it returns the ASCII value of the byte subtracted by the ASCII value of `0`, and `true`.
 
The `IsValidISBN()` function starts be defining variables for iterating the characters in the input `string`.

The ways to iterate characters are by Unicode runes, or by each letter being a string, or by each letter being a byte.
The runes are from [`range`][range] on a string, the strings from [`Split()`][split], and the bytes from indexing into the string.
Another way to iterate runes is to convert the string to a rune [`slice`][slice] and `range` on it.
The difference between ranging on a rune slice vs ranging on a string is that the index returned from a string is the position of the next rune in bytes,
not which rune it is.
For example, if the first unicode character is two bytes,
then the second unicode character index will be 2 when ranging on a string and 1 when ranging on a rune slice.
As of the time of this writing we can iterate bytes, since all of the characters are ASCII.

The character is picked from the input `string` by its index.
It is then checked to see if it is a digit, which is the expected and usual value.
It might be expected to check the input for invalid values first, thus aligning the ["happy path"][happy-path] to the left edge,
but here it is simpler and more efficient to check for the valid and expected value first.
Most characters will usually be a digit, and the happy path is indented only once, so it is not too onerous.

If the character is a digit, then it is multiplied by its position and is added to the sum,
and the position is decremented.

Otherwise, if the position is `1`  and the character is an `X`, then `10` is added to the sum,
and the position is decremented.

If neither of those two conditions apply, then, if the character is not a dash,
it is an illegal character and the function immediately returns `false`.
If the character is a dash, then the loop continues without decrementing the position.

After the loop finishes, the function returns whether the position is at the expected value of `0`,
and if the sum is evenly divisible by 11.

[bytes]: https://pkg.go.dev/bytes 
[ascii]: https://www.asciitable.com/
[named-return-values]: https://yourbasic.org/golang/named-return-values-parameters/
[zero-values]: https://yourbasic.org/golang/default-zero-value/
[range]: https://gobyexample.com/range
[split]: https://pkg.go.dev/strings#Split
[slice]: https://gobyexample.com/slices
[happy-path]: https://en.wikipedia.org/wiki/Happy_path
