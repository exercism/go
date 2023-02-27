# Bit-shifting

```go
// Package grains is a small library for determing the squares of a number.
package grains

import (
	"errors"
)

// Square returns the number of grains on the square for the number passed in.
func Square(num int) (uint64, error) {
	if num < 1 || num > 64 {
		return 0, errors.New("square number must be 1 through 64")
	}

	return 1 << (num - 1), nil
}

// Total return the sum of the squares from 1 to 64, which is the number of squares on a chess board.
func Total() uint64 {
	return 1<<64 - 1
}
```

Instead of using math to calculate the number of grains on a square, you can set a bit in the correct position of a [`uint64`][uint64].

To understand how this works, consider just two squares that are represented in binary bits as `00`.

You use the [left-shift operator][left-shift-operator] to set `1` at the position needed to make the correct decimal value.
- To set the one grain on Square One you shift `1` for `0` positions to the left.
So, if `n` is `1` for square One, you subtract `n` by `1` to get `0`, which will not move it any positions to the left.
The result is binary `01`, which is decimal `1`.
- To set the two grains on Square Two you shift `1` for `1` position to the left.
So, if `n` is `2` for square Two, you subtract `n` by `1` to get `1`, which will move it `1` position to the left.
The result is binary `10`, which is decimal `2`.

| Square  | Shift Left By | Binary Value | Decimal Value |
| ------- | ------------- | ------------ | ------------- |
|       1 |             0 |         0001 |             1 |
|       2 |             1 |         0010 |             2 |
|       3 |             2 |         0100 |             4 |
|       4 |             3 |         1000 |             8 |

For `total` we want all of the 64 bits set to `1` to get the sum of grains on all sixty-four squares.
The easy way to do this is to set the 65th bit to `1` and then subtract `1`.
To go back to our two-square example, if we can grow to three squares, then we can shift `1` two positions to the left for binary `100`,
which is decimal `4`.
By subtracting `1` we get `3`, which is the total amount of grains on the two squares.

| Square  | Binary Value | Decimal Value |
| ------- | ------------ | ------------- |
|       3 |         0100 |             4 |

| Square  | Sum Binary Value | Sum Decimal Value |
| ------- | ---------------- | ----------------- |
|       1 |             0001 |                 1 |
|       2 |             0011 |                 3 |


```exercism/note
A shortcut would be to use the [`math.MaxUint64`](https://cs.opensource.google/go/go/+/refs/tags/go1.19.4:src/math/const.go;l=39)
constant, which is defined as `1<<64 - 1`.
```

[uint64]: https://pkg.go.dev/builtin#uint64
[left-shift-operator]: https://www.golangprograms.com/bitwise-operators-in-go-programming-language.html
[maxuint64]: https://cs.opensource.google/go/go/+/refs/tags/go1.19.4:src/math/const.go;l=39
