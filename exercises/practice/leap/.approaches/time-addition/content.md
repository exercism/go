# `time.Add()`

```go
// Package leap is a small library for determining if the passed in year is a leap year.
package leap

import "time"

// IsLeapYear returns if the passed in year is a leap year.
func IsLeapYear(year int) bool {
	return time.Date(year, time.February, 28, 0, 0, 0, 0, time.UTC).Add(time.Duration(time.Hour*24)).Day() == 29
}
```

This approach may be considered a "cheat" for this exercise.
By adding a day to February 28th for the year, you can see if the new day is the 29th or the 1st.
If it is the 29th, then the year is a leap year.
This is done by using the [time.Add()][time-add] and [time.Day()][time-day] functions.

[time-add]: https://pkg.go.dev/time#Time.Add
[time-day]: https://pkg.go.dev/time#Time.Day
