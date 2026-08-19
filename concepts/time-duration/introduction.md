# Introduction

In Go, time-based calculations are handled using the [`time`][time-package]  package.  
Two key types used in this exercise are:

- [`time.Time`][time]: represents a specific moment in time (a timestamp)
- [`time.Duration`][duration]: represents the amount of time elapsed between two moments

---

## `time.Time` vs `time.Duration`

A `time.Time` represents an absolute point in time, such as:

```go
t := time.Now()
```

A `time.Duration` represents the difference between two points in time:
```go
start := time.Now()
end := start.Add(5 * time.Minute)

duration := end.Sub(start) // => 5m0s
```

In short:
- `time.Time` → “when something happens”
- `time.Duration` → “how long it lasts”

## Measuring Time with `time.Since`
To measure how long something takes (e.g., execution time), you can use `time.Since`[since]
```go
start := time.Now()

// some work here

elapsed := time.Since(start)
fmt.Println(elapsed)
```

## Adding and Subtracting Time
You can manipulate timestamps using:
- [`Time.Add()`][add] → add or subtract a duration
- [`Time.Sub()`][sub] → calculate the difference between two times

## Time Units
Go provides built-in constants for time units:

| Unit        | Constant             |
| ----------- |----------------------|
|Nanosecond   | `time.Nanosecond`    |
|Microsecond  | `time.Microsecond `  |
|Millisecond  | `time.Millisecond`   |
|Second       | `time.Second`        |
|Minute       | `time.Minute`        |
|Hour         | `time.Hour`          |

You can combine them to build durations:
```go
d := 1*time.Hour + 5*time.Minute + 19*time.Second + 234*time.Millisecond
```

You can also perform calculations:
```go
seconds := d.Seconds()   // float64
minutes := d.Minutes()
```

## Working with Durations
Durations can be:
- Added to a `time.Time`
- Subtracted from a `time.Time`
- Compared or converted to different units

```go
start := time.Now()
finish := start.Add(30 * time.Second)

duration := finish.Sub(start) // => 30s
```

## Other resources
- [Go by Example: Time](https://gobyexample.com/time)
- [Go by Example: Time Formatting / Parsing](https://gobyexample.com/time-formatting-parsing)

[time-package]: https://pkg.go.dev/time
[time]: https://pkg.go.dev/time#Time
[duration]: https://pkg.go.dev/time#Duration
[since]: https://pkg.go.dev/time#Since
[add]: https://pkg.go.dev/time#Time.Add
[sub]: https://pkg.go.dev/time#Time.Sub
[now]: https://pkg.go.dev/time#Now

