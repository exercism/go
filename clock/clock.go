package clock

import "fmt"

const (
	MinutesPerHour = 60
	MinutesPerDay  = 1440
)

type Clock int

func New(h, m int) Clock {
	c := Clock((h*MinutesPerHour + m) % MinutesPerDay)
	if c < 0 {
		c += MinutesPerDay
	}
	return c
}

func (c Clock) Add(m int) Clock {
	return New(0, int(c)+m)
}

func (c Clock) Subtract(m int) Clock {
	return c.Add(-m) //New(0, int(c)-m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/MinutesPerHour, c%MinutesPerHour)
}
