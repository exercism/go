# Chain of boolean expressions

```go
// Package leap is a small library for determing if the passed in year is a leap year.
package leap

// IsLeapYear returns if the passed in year is a leap year.
func IsLeapYear(year int) bool {
	return  year%4 == 0 && (year%100 != 0 || year%400 == 0)
}
```

The first boolean expression uses the [modulus operator][modulus-operator] to check if the year is evenly divided by `4`.
- If the year is not evenly divisible by `4`, then the chain will "short circuit" due to the next operator being a [logical AND][logical-and] (`&&`),
and will return `false`.
- If the year _is_ evenly divisible by `4`, then the year is checked to _not_ be evenly divisible by `100`.
- If the year is not evenly divisible by `100`, then the expression is `true` and the chain will "short-circuit" to return `true`,
since the next operator is a [logical OR][logical-or] (`||`).
- If the year _is_ evenly divisible by `100`, then the expression is `false`, and the returned value from the chain will be if the year is evenly divisible by `400`.

| year | year % 4 == 0 | year % 100 != 0 | year % 400 == 0 | is leap year |
| ---- | ------------- | --------------- | --------------- | ------------ |
| 2020 |          true |            true |   not evaluated |         true |
| 2019 |         false |   not evaluated |   not evaluated |        false |
| 2000 |          true |           false |            true |         true |
| 1900 |          true |           false |           false |        false |

[modulus-operator]: https://golangbyexample.com/remainder-modulus-go-golang/
[logical-and]: https://www.tutorialkart.com/golang-tutorial/golang-and/
[logical-or]: https://www.tutorialkart.com/golang-tutorial/golang-or/
