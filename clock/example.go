package clock

import "fmt"

type Clock int

func New(h, m int) Clock {
	c := Clock((h*60 + m) % 1440)
	if c < 0 {
		c += 1440
	}
	return c
}

func (c Clock) Add(m int) Clock {
	c = (c + Clock(m)) % 1440
	if c < 0 {
		c += 1440
	}
	return c
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}
