# Maximum of two checks

```go
// Package leap is a small library for determing if the passed in year is a leap year.
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

This approach uses a maximum of two checks to determine if a year is a leap year.

It starts by testing the outlier condition of the year being evenly divisible by `100`.
It does this by using the [modulus operator][modulus-operator].
If the year is evenly divisible by `100`, then the expression is `true`, and the function returns if the year is evenly divisible by `400`.
If the year is _not_ evenly divisible by `100`, then the expression is `false`, and the function returns if the year is evenly divisible by `4`.

| year | year % 100 == 0 | year % 400 == 0 | year % 4 == 0  | is leap year |
| ---- | --------------- | --------------- | -------------- | ------------ |
| 2020 |           false |   not evaluated |           true |        true  |
| 2019 |           false |   not evaluated |          false |       false  |
| 2000 |           true  |            true |  not evaluated |        true  |
| 1900 |           true  |           false |  not evaluated |        false |

Although it uses a maximum of only two checks, this approach tests an outlier condition of the year being evenly divisible by `100` first,
which is less likely than the year being evenly divisible by `4`.

[modulus-operator]: https://golangbyexample.com/remainder-modulus-go-golang/
