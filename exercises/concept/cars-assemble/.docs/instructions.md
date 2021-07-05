# Instructions

In this exercise you'll be writing code to analyse the production of an
assembly line in a car factory. The assembly line's speed can range from `0`
(off) to `10` (maximum).

At its default speed (`1`), `221` cars are produced each hour. In principle,
the production increases linearly. So with the speed set to `4`, it should
produce `4 * 221 = 884` cars per hour. However, higher speeds increase the
likelihood that faulty cars are produced, which then have to be discarded. The
following table shows how speed influences the success rate:

- `0`: 0% success rate.
- `1` - `4`: 100% success rate.
- `5` - `8`: 90% success rate.
- `9` - `10`: 77% success rate.

> Note that the calculation of the success rate has been provided already.

You have two tasks:

## 1. Calculate the production rate per hour

Implement a function to calculate the assembly line's production rate per hour.

```go
rate := CalculateProductionRatePerHour(7)
fmt.Println(rate)
// Output: 10829.0
```

> Note that the value returned is of type `float64`

## 2. Calculate the number of working items produced per minute

Implement a function to calculate how many cars are produced each minute:

```go
rate := CalculateProductionRatePerMinute(5)
fmt.Println(rate)
// Output: 92
```

> Note that the value returned is of type `int`.
