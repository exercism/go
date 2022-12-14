# Scrub out the spaces and test if the input is an integer first

```go
// Package luhn is a small library for returning if a phrase is valid according to the Luhn algorithm.
package luhn

import (
	"strconv"
	"strings"
)

func Valid(num string) bool {
	num = strings.ReplaceAll(num, " ", "")
	if _, err := strconv.Atoi(num); err != nil {
		return false
	}
	idx := len(num) - 1
	total := 0
	pos := 0
	for i := idx; i > -1; i-- {
		digit := int(num[i] - '0')
		if pos%2 == 0 {
			total += digit
		} else {
			switch digit {
			case 1, 2, 3, 4:
				total += digit << 1
			case 5, 6, 7, 8:
				total += (digit << 1) - 9
			case 9:
				total += digit
			}
		}
		pos++
	}
	return pos > 1 && total%10 == 0
}
```

This approach starts with removing all spaces in the input by calling the [`strings`][strings] [`ReplaceAll()`][replaceall] function.
The [`strconv`][strconv] [`Atoi()`][atoi] function is then used to see if the input string can be converted to an integer.
If not, the function returns `false`.
Some variables are then initialized for iterating backward (i.e. from right to left) through the input string.

The ways to iterate characters are by Unicode runes, or by each letter being a string, or by each letter being a byte.
The runes are from [`range`][range] on a string, the strings from [`Split()`][split], and the bytes from indexing into the string.
Another way to iterate runes is to convert the string to a rune [`slice`][slice] and `range` on it.
The difference between ranging on a rune slice vs ranging on a string is that the index returned from a string is the position of the next rune in bytes,
not which rune it is.
For example, if iterating from left to right and the first unicode character is two bytes,
then the second unicode character index will be 2 when ranging on a string and 1 when ranging on a rune slice.
As of the time of this writing we can iterate bytes, since all of the characters are ASCII.

Iterating backward is a way to keep us from having to reverse the input string.

In the loop, the character is converted to an `int` after subtracting its [ASCII][ascii] value by the ASCII value of `'0'`,
for example `'3'` - `'0'` = `3`.

The `pos` variable is then checked to see if it is evenly divisble by `0`.
If so, it adds the digit value to the total.

If not, it adds double the digit value to the total if the digit value is less than `5`.
This can be done by shifting all of the digit bits one position to the left, or by multiplying by `2`.
For example, `3` in binary is `11`, By shifting all of the bits once to the left it becomes `110`,
which is `6` in binary (`100` is `4` plus `10` is `2` = `110` is `6`.)

Otherwise, if the digit value is greater than `4` and less than `9`, then `9` is subtracted from the doubled digit value
and then added to the total.

If the digit value is `9`, then it is added to the total.

The `pos` variable is incremented to keep track of the position, and the loop either exits or goes on to the next digit.

After the loop exits, the functon returns if the position is greater than `1` and the total is evenly divisible by `10`.

[strings]: https://pkg.go.dev/strings
[replaceall]: https://pkg.go.dev/strings#ReplaceAll
[strconv]: https://pkg.go.dev/strconv
[atoi]: https://pkg.go.dev/strconv#Atoi
[range]: https://gobyexample.com/range
[split]: https://pkg.go.dev/strings#Split
[slice]: https://gobyexample.com/slices
[ascii]: https://www.asciitable.com/
