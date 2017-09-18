package gigasecond

import "time"

// AddGigasecond returns time t plus one gigasecond.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
