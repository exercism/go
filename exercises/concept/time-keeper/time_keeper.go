package time_keeper

import (
	"time"
)

// GetDuration - Get the difference between start and finish.
func GetDuration(start, finish time.Time) time.Duration {
	panic("Please implement GetDuration")
}

// GetDurationInDifferentTimezones - Get the difference between times from different time zones.
func GetDurationInDifferentTimezones(start, finish time.Time) time.Duration {
	panic("Please implement GetDurationInDifferentTimezones")
}

// ParseRun - Parse a duration like "10h5m19.234s".
func ParseRun(input string) (time.Duration, error) {
	panic("Please implement ParseRun")
}

// FixStartTime - Recalculate correct start time.
func FixStartTime(incorrectStart time.Time, offset time.Duration) time.Time {
	panic("Please implement FixStartTime")
}

// IsValidDetection - Check whether a detection is within an allowed time window after start.
func IsValidDetection(start time.Time, from, to time.Duration, detection time.Time) bool {
	panic("Please implement IsValidDetection")
}

// FormatResult - Format leaderboard output.
func FormatResult(result, leader time.Duration) string {
	panic("Please implement FormatResult")
}
