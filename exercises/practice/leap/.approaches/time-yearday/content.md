# `time.YearDay()`

```go
// Package leap is a small library for determing if the passed in year is a leap year.
package leap

import "time"

// IsLeapYear returns if the passed in year is a leap year.
func IsLeapYear(year int) bool {
	return time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay() == 366
}
```

This approach may be considered a "cheat" for this exercise.
The [`time.YearDay()`][time-yearday] function is used to see if December 31 for the year is the 366th day of the year.
If there are 366 days in the year, then the function returns `true`.

[time-yearday]: https://pkg.go.dev/time#Time.YearDay
