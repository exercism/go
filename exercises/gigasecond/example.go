package gigasecond

import "time"

const testVersion = 4

// AddGigasecond returns time t plus one gigasecond.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
