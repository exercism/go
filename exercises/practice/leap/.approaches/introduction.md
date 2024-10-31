# Introduction

There are various idiomatic approaches to solve Leap.
You can use a chain of boolean expressions to test the conditions.
Or you can use a maximum of two checks.

## General guidance

The key to solving Leap is to know if the year is evenly divisible by `4`, `100` and `400`.
For determining that, you will use the [modulus operator][modulus-operator].

## Approach: Chain of Boolean expressions

```go
// Package leap is a small library for determining if the passed in year is a leap year.
package leap

// IsLeapYear returns if the passed in year is a leap year.
func IsLeapYear(year int) bool {
	return  year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
```

For more information, check the [Boolean chain approach][approach-boolean-chain].

## Approach: Maximum of two checks

```go
// Package leap is a small library for determining if the passed in year is a leap year.
package leap

// IsLeapYear returns if the passed in year is a leap year.
func IsLeapYear(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	} else {
		return year%4 == 0
	}
}
```

For more information, check the [maximum of two checks approach][approach-max-two-checks].

## Other approaches

Besides the aforementioned, idiomatic approaches, you could also approach the exercise as follows:

### `time.Add()` approach:

Add a day to February 28th for the year and see if the new day is the 29th. For more information, see the [`time.Add()` approach][approach-time-add].

### `time.YearDay()` approach:

See if the `YearDay()` for December 31 is 366. For more information, see the [`time.YearDay()` approach][approach-time-yearday].

## Which approach to use?

The chain of boolean expressions _should_ be most efficient, as it proceeds from the most likely to least likely conditions.
However,  the maximum of two checks approach benched faster, even though it starts from a less likely condition.
Using `time.Add()` or `time.YearDay()` may be considered "cheats" for the exercise,
but are both much slower than the other approaches.

For more information, check the [Performance article][article-performance].

[modulus-operator]: https://golangbyexample.com/remainder-modulus-go-golang/
[approach-boolean-chain]: https://exercism.org/tracks/go/exercises/leap/approaches/boolean-chain
[approach-max-two-checks]: https://exercism.org/tracks/go/exercises/leap/approaches/max-two-checks
[approach-time-add]: https://exercism.org/tracks/go/exercises/leap/approaches/time-addition
[approach-time-yearday]: https://exercism.org/tracks/go/exercises/leap/approaches/time-yearday
[article-performance]: https://exercism.org/tracks/go/exercises/leap/articles/performance
