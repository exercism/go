# `Clock` as an `int`

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
	return fmt.Sprintf("%02d:%02d", clock/hourMinutes, clock%hourMinutes)
}
```

The approach imports the `fmt` package so it can use the [`fmt.Sprintf()`][sprintf] function in the `String()` method.

This approach starts by typing `Clock` as an `int`.

It then defines a couple of [constants][const] with meaningful names for `1440` and `60`,
instead of using the literals as [magic numbers][magic-numbers].

The `normalize()` helper function is defined to properly handle rollover for when adding or subtracting minutes crosses into the next or previous day(s).
There are other ways to handle rollover, such as the following expression:

```go
((hours*hourMinutes+minutes)%dayminutes + dayMinutes) % dayMinutes
```

Since `normalize()` returns a `Clock`, it can be used by the `New` function as well as the `Add` and `Subtract` [methods][methods].

```exercism/note
Note that, in the `Add` and `Subtract`methods, the received `Clock` argument must be converted to an `int` to be
added to or subtracted by the `int` minutes, even though the underlying type of `Clock` is an `int`.
This can be helpful if two `int` values are two different types, such as a meter and a foot.
The compiler will not allow adding a meter type to a foot type,
even if both are an `int`, unless one of them is converted to the other.
For an example of why this is important, see an article on how a $125 million spacecraft crashed because of a 
[mix-up between English and metric units of measure](https://www.simscale.com/blog/nasa-mars-climate-orbiter-metric/).
```

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
