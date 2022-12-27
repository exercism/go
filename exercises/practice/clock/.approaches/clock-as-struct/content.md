# `Clock` as a struct

```go
// Package Clock is a small library for keeping track of time.
package clock

import (
	"fmt"
)

const (
	dayMinutes  = 1440
	hourMinutes = 60
)

// Clock type
type Clock struct {
	minutes int
}

func normalize(minutes int) Clock {
	if minutes < 0 {
		return Clock{dayMinutes + (minutes % dayMinutes)}
	} else if minutes >= dayMinutes {
		return Clock{minutes % dayMinutes}
	}
	return Clock{minutes}
}

// New returns a Clock that represents the hour and minutes.
func New(hour, minutes int) Clock {
	return normalize(hour*hourMinutes + minutes)
}

// Subtract returns a Clock that represents the Clock plus minutes.
func (c Clock) Add(minutes int) Clock {
	return normalize(c.minutes + minutes)
}

// Add returns a Clock that represents the Clock minus minutes.
func (c Clock) Subtract(minutes int) Clock {
	return normalize(c.minutes - minutes)
}

// String returns the Clock formatted into hours:minutes, e.g. 01:02.
func (c *Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/hourMinutes, c.minutes%hourMinutes)
}
```

The approach imports the `fmt` package so it can use the [`fmt.Sprintf()`][sprintf] function in the `String()` method.

This approach starts by typing `Clock` as a struct with an `int` field.

It then defines a couple of [constants][const] with meaningful names for `1440` and `60`,
instead of using the literals as [magic numbers][magic-numbers].

The `normalize()` helper function is defined to properly handle rollover for when adding or subtracting minutes crosses into the next or previous day(s).
There are other ways to handle rollover, such as the following expression:

```go
((hours*hourMinutes+minutes)%dayminutes + dayMinutes) % dayMinutes
```

Since `normalize()` returns a `Clock`, it can be used by the `New` function as well as the `Add` and `Subtract` [methods][methods].

The `String()` function uses [`fmt.Sprintf()`][sprintf] to format the hour and minutes with leading zeros.
If the amount of minutes is `69`, then the hour is `69` divided by `60` (1),
and the minutes is the remainder of `69` divided by `60` (9), so the string for `69` minutes would be `01:09`.
The format specifier `%02d` says to pad a base-10 number with zeros for up to two places.
If one place is occupied by a digit, then the extra place is padded with a zero.
More info can be found at a [format cheat sheet][format-cheat-sheet].

[sprintf]: https://pkg.go.dev/fmt#Sprintf
[const]: https://go.dev/tour/basics/15
[magic-numbers]: https://en.wikipedia.org/wiki/Magic_number_(programming)
[methods]: https://go.dev/tour/methods/1
[format-cheat-sheet]: https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/
