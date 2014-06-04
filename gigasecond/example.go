package gigasecond

import "time"

// A gigasecond is 10^9 seconds.
const Gigasecond time.Duration = 1e9 * time.Second

// AddGigasecond returns the time t + Gigasecond.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
