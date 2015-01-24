package gigasecond

import "time"

const TestVersion = 2

// AddGigasecond returns time t plus one gigasecond.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}

var Birthday = time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
