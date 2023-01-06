# Introduction

There are various idiomatic approaches to solve Grains.
You can use [`math.Pow()`][math-pow] to calculate the number on grains on a square, and
[`big` exponentiation][big-exponentiation] to calculate the total.
Or you can use bit-shifting.

## General guidance

The key to solving Grains is to focus on each square having double the amount of grains as the square before it.
This means that the amount of grains grows exponentially.
The first square has one grain, which is `2` to the power of `0`.
The second square has two grains, which is `2` to the power of `1`.
The third square has four grains, which is `2` to the power of `2`.
You can see that the exponent, or power, that `2` is raised by is always one less than the square number.

| Square | Power | Value                   |
| ------ | ----- | ----------------------- |
| 1      | 0     | 2 to the power of 0 = 1 |
| 2      | 1     | 2 to the power of 1 = 2 |
| 3      | 2     | 2 to the power of 2 = 4 |
| 4      | 3     | 2 to the power of 4 = 8 |

You can use the [`big`][big] package and its [Int][bigint] type to support numbers above [`math.MaxUint64`][maxuint64].

## Approach: `math.Pow()` and `big` Exponentiation

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

For more information, check the [`math.Pow()` and `big` exponentiation approach][approach-math-pow-big-exponentiation].

## Approach: Bit-shifting

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

For more information, check the [Bit-shifting approach][approach-bit-shifting].

## Which approach to use?

The bit shifting approach is much faster than using `math.Pow()` and `big` exponentiation.

To compare the performance of the approaches, check out the [Performance article][article-performance].

[math-pow]: https://pkg.go.dev/math#Pow
[big-exponentiation]: https://pkg.go.dev/math/big#Int.Exp
[big]: https://pkg.go.dev/math/big
[bigint]: https://pkg.go.dev/math/big#Int
[maxuint64]: https://www.includehelp.com/golang/math-maxuint64-constant-with-examples.aspx
[approach-math-pow-big-exponentiation]: https://exercism.org/tracks/go/exercises/grains/approaches/math-pow-big-exponentiation
[approach-bit-shifting]: https://exercism.org/tracks/go/exercises/grains/approaches/bit-shifting
[article-performance]: https://exercism.org/tracks/go/exercises/grains/articles/performance
