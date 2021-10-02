# Instructions

In this exercise you'll be writing code to analyze the production of an assembly line in a car factory.
The assembly line's speed can range from `0` (off) to `10` (maximum).

At its default speed (`1`), `221` cars are produced each hour.
In principle, the production increases linearly.
So with the speed set to `4`, it should produce `4 * 221 = 884` cars per hour.
However, higher speeds increase the likelihood that faulty cars are produced, which then have to be discarded.

## 1. Calculate the success rate

Implement a function (`SuccessRate`) to calculate the ratio of an item being created without error for a given speed.
The following table shows how speed influences the success rate:

- `0`: 0% success rate.
- `1` - `4`: 100% success rate.
- `5` - `8`: 90% success rate.
- `9` - `10`: 77% success rate.

```go
rate := SuccessRate(6)
fmt.Println(rate)
// Output: 0.9
```

## 2. Calculate the production rate per hour

Implement a function to calculate the assembly line's production rate per hour.

```go
rate := CalculateProductionRatePerHour(7)
fmt.Println(rate)
// Output: 1392.3
```

> Note that the value returned is of type `float64`.

## 3. Calculate the number of working items produced per minute

Implement a function to calculate how many cars are produced each minute:

```go
rate := CalculateProductionRatePerMinute(5)
fmt.Println(rate)
// Output: 16
```

> Note that the value returned is of type `int`.
