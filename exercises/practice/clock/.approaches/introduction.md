# Introduction

There are at least a couple of idiomatic ways to solve Clock.
One approach is to type `Clock` as an `int`
Another approach is to type `Clock` as a struct with an `int` field.

## General guidance

Something to consider is how the logic for calculation may possibly be simplified by storing only minutes in the Clock type.
That leaves [separating the concern][separating the concern] of hours and minutes only for the `String()` function,
which is the only place that cares about it.

Instead of using `60`, `24` or `1440` literals as [magic numbers][magic numbers],
it might be considered to define them as [constants][const] with informative names.

Another concern is to not duplicate code between `New`, `Add` and `Subtract`,
thus keeping the code [DRY](https://en.wikipedia.org/wiki/Don%27t_repeat_yourself).

Finally, for `Add` and `Subtract`, consider returning a new instance of `Clock` instead of mutating the existing `Clock`.
For the benefits of immutability in general, see [this article][immutability-benefits].

## Approach: `Clock` as an `int`

```go
// Package Clock is a small library for keeping track of time.
package clock

import "fmt"

// Clock type
type Clock int

const (
	dayMinutes  = 1440
	hourMinutes = 60
)

func normalize(minutes int) Clock {
	if minutes < 0 {
		return Clock(dayMinutes + (minutes % dayMinutes))
	} else if minutes >= dayMinutes {
		return Clock(minutes % dayMinutes)
	}
	return Clock(minutes)
}

// New returns a Clock that represents the hour and minutes.
func New(hour, minutes int) Clock {
	return normalize(hour*hourMinutes + minutes)
}

// Add returns a Clock that represents the Clock plus minutes.
func (clock Clock) Add(minutes int) Clock {
	return normalize(int(clock) + minutes)
}

// Subtract returns a Clock that represents the Clock minus minutes.
func (clock Clock) Subtract(minutes int) Clock {
	return normalize(int(clock) - minutes)
}

// String returns the Clock formatted into hours:minutes, e.g. 01:02.
func (clock Clock) String() string {
	return fmt.Sprintf("%02d:%02d", clock/60, clock%60)
}
```

For more information, check the [`Clock` as an `int` approach][approach-clock-as-int].

## Approach: `Clock` as a struct

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

// String returns the Clock formatted into hours:minutes, e.g. 01:02.
func (c *Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.minutes/60, c.minutes%60)
}

// Subtract returns a Clock that represents the Clock plus minutes.
func (c Clock) Add(minutes int) Clock {
	return normalize(c.minutes + minutes)
}

// Add returns a Clock that represents the Clock minus minutes.
func (c Clock) Subtract(minutes int) Clock {
	return normalize(c.minutes - minutes)
}
```

For more information, check the [`Clock` as a struct approach][approach-clock-as-struct].

## Which approach to use?

Both approaches are similar in performance, so which to use could be a matter of style preference.
To compare performance of the approaches, check the [Performance article][article-performance].

[separating the concern]: https://en.wikipedia.org/wiki/Separation_of_concerns
[magic numbers]: https://en.wikipedia.org/wiki/Magic_number_(programming)
[const]: https://go.dev/tour/basics/15
[DRY]: https://en.wikipedia.org/wiki/Don%27t_repeat_yourself
[immutability-benefits]: https://hackernoon.com/5-benefits-of-immutable-objects-worth-considering-for-your-next-project-f98e7e85b6ac
[approach-clock-as-int]: https://exercism.org/tracks/go/exercises/clock/approaches/clock-as-int
[approach-clock-as-struct]: https://exercism.org/tracks/go/exercises/clock/approaches/clock-as-struct
[article-performance]: https://exercism.org/tracks/go/exercises/clock/articles/performance
