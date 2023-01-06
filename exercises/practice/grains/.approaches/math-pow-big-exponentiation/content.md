## `math.Pow()` and `big` exponentiation

```go
package grains

import (
	"errors"
	"math"
	"math/big"
)

// Square returns the number of grains on the square for the number passed in.
func Square(num int) (uint64, error) {
	if num < 1 || num > 64 {
		return 0, errors.New("square number must be 1 through 64")
	}

	return uint64(math.Pow(2, float64(num-1))), nil
}

// Total return the sum of the squares from 1 to 64, which is the number of squares on a chess board.
func Total() uint64 {
	var num, exp = big.NewInt(2), big.NewInt(64)
	return num.Exp(num, exp, nil).Sub(num, big.NewInt(1)).Uint64()
}
```

Go uses [`math.Pow()`][math-pow] to raise a number by a certain exponent.
[`big.Int.Exp()`][exp] can also be used to raise a number by an exponent to work with a [`big.Int`][bigint].

Exponentiation is nicely suited to the problem, since we start with one grain and keep doubling the number of grains on each successive square.
`1` grain is `math.Pow(2, 0)`, `2` grains is `math.Pow(2, 1)`, `4` is `math.Pow(2, 2)`, and so on.

So, to get the right exponent, we subtract `1` from the square number `num`.

The easiest way to get `total` is to get the value for an imaginary 65th square,
and then subtract `1` from it.
To understand how that works, consider a board that has only two squares.
If we could grow the board to three squares, then we could get the number of grains on the imaginary third square,
which would be `4.`
You could then subtract `4` by `1` to get `3`, which is the number of grains on the first square (`1`) and the second square (`2`).
Remembering that the exponent must be one less than the square you want,
you would want to call `math.Pow(2, 64)` to get the number of grains on the imaginary 65th square.
However, that won't work, because the resulting number is too large.
So, you can use [`big.NewInt`][newint] to create the numbers to pass to [`Exp()`][exp].
The output from `Exp()` is then used to call the [`Sub()`][sub] function for subtracting that value by `1`,
and its output is used for calling [`Uint64()`][uint64],
which converts the big `Int` value to the `uint64` value returned for all `64` squares.

[math-pow]: https://pkg.go.dev/math#Pow
[exp]: https://pkg.go.dev/math/big#Int.Exp
[bigint]: https://pkg.go.dev/math/big#Int
[newint]: https://pkg.go.dev/math/big#NewInt
[sub]: https://pkg.go.dev/math/big#Int.Sub
[uint64]: https://pkg.go.dev/math/big#Int.Uint64
