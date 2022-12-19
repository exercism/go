# Validate as you go

```go
// Package luhn is a small library for returning if a phrase is valid according to the Luhn algorithm.
package luhn

func Valid(num string) bool {
	total := 0
	pos := 0
	for i := len(num) - 1; i > -1; i-- {
		char := num[i]
		if char == ' ' {
			continue
		}
		if char < '0' || char > '9' {
			return false
		}
		digit := int(char - '0')
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

This approach starts by initializing some variables for iterating backward (i.e. from right to left) through the input string.

The ways to iterate characters are by Unicode runes, or by each letter being a string, or by each letter being a byte.
The runes are from [`range`][range] on a string, the strings from [`Split()`][split], and the bytes from indexing into the string.
Another way to iterate runes is to convert the string to a rune [`slice`][slice] and `range` on it.
The difference between ranging on a rune slice vs ranging on a string is that the index returned from a string is the position of the next rune in bytes,
not which rune it is.
For example, if iterating from left to right and the first unicode character is two bytes,
then the second unicode character index will be 2 when ranging on a string and 1 when ranging on a rune slice.
As of the time of this writing we can iterate bytes, since all of the characters are ASCII.

Iterating backward is a way to keep us from having to reverse the input string.

In the loop, the character is checked to see if it is a space.
If so, [continue][continue] is used to go to the next iteration without incrementing the `pos` variable,
thus ignoring the space and indicating the same position.

If it is not a space, then the character is tested to be a digit character.
If not, the function returns `false` for having an illegal character.

The character is converted to an `int` after subtracting its [ASCII][ascii] value by the ASCII value of `'0'`,
for example `'3'` - `'0'` = `3`.

The `pos` variable is then checked to see if the position is evenly divisible by `0`.
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

# Optimization

There is an interesting optimzation for the validate as you go approach that can cut the time by more than half.
It requires two modifications to be made in tandem.
Either one by itself does not reduce the time so dramatically.
It was benchmarked on version `1.19.4 windows/amd`.

The whole modified code is here

```go
// Package luhn is a small library for returning if a phrase is valid according to the Luhn algorithm.
package luhn

func Valid(num string) bool {
    var total, pos int
        for i := len(num) - 1; i > -1; i-- {
            char := num[i]
            if char == ' ' {
                continue
            }
            if char < 48 || char > 57 {
                return false
            }
            digit := int(char - 48)
            total += digit + pos%2*(digit+digit/5)
            pos++
        }
        return pos > 1 && total%10 == 0
}
```

The first modification is simple and changes 

```go
total := 0
pos := 0
```

for 

```
var total, pos int
```

which defines the same variables with their defaulrt [zero values][zero-values].

The other change is more of a mathematical way to avoid the conditional logic for adding the value.
The `if` statement and entire `switch` statement can be replaced with 

```go
total += digit + pos%2*(digit+digit/5)
```

The first mathematical choice is that when the position number is even,
the number multiplies its "doubled" value by `0`, so evenly-positioned numbers only add themselves.
An oddly-positioned number will multiply its "doubled" value by `1`.

The second mathemtical choice is that, instead of possibly subtracting "down" by 9, the expression always adds "up".
For higher numbers that would have been subtracted by `9`, the result is the number desired plus `10`.
Since the total value is checked by seeing if it is evenly divisible by `10`,
the extra `10` in the higher-number calculations is effectively factored out.

To understand how this works, you can see in the following table for `digit + (digit + digit/5)`

| Value   | Desired Value | Actual Value | How                            |
| ------- | ------------- | ------------ | -----------------------------  |
|       1 |             2 |            2 |  1 + (1 + 1/5(0)) = 2)  |
|       2 |             4 |            4 |  2 + (2 + 2/5(0)) = 4)  |
|       3 |             6 |            6 |  3 + (3 + 3/5(0)) = 6)  |
|       4 |             8 |            8 |  4 + (4 + 4/5(0)) = 8)  |
|       5 |             1 |           11 |  5 + (5 + 5/5(1)) = 11) |
|       6 |             3 |           13 |  6 + (6 + 6/5(1)) = 13) |
|       7 |             5 |           15 |  7 + (7 + 7/5(1)) = 15) |
|       8 |             8 |           17 |  8 + (8 + 8/5(1)) = 17) |
|       9 |             9 |           19 |  9 + (9 + 9/5(1)) = 19) |


The benchmarks were as follows:

```
unoptimizied
BenchmarkValid-12    	 4685019	       251.5 ns/op	       0 B/op	       0 allocs/op

optimized
BenchmarkValid-12    	10885153	       106.8 ns/op	       0 B/op	       0 allocs/op
```

The optimized solution is much faster, but it is not as clear as to how it is conforming to the Luhn algorithm,
unless one readily understands the math.

[strings]: https://pkg.go.dev/strings
[replaceall]: https://pkg.go.dev/strings#ReplaceAll
[strconv]: https://pkg.go.dev/strconv
[atoi]: https://pkg.go.dev/strconv#Atoi
[range]: https://gobyexample.com/range
[split]: https://pkg.go.dev/strings#Split
[slice]: https://gobyexample.com/slices
[continue]: https://www.tutorialspoint.com/go/go_continue_statement.htm
[ascii]: https://www.asciitable.com/
[zero-values]: https://yourbasic.org/golang/default-zero-value/
