package time_keeper

import (
	"fmt"
	"time"
)

// GetDuration - Get the difference between start and finish.
func GetDuration(start, finish time.Time) time.Duration {
	return finish.Sub(start)
}

// GetDurationInDifferentTimezones - Get the difference between times from different time zones.
func GetDurationInDifferentTimezones(start, finish time.Time) time.Duration {
	return finish.Sub(start)
}

// ParseRun - Parse a duration like "10h5m19.234s".
func ParseRun(input string) (time.Duration, error) {
	return time.ParseDuration(input)
}

// FixStartTime - Recalculate correct start time.
func FixStartTime(incorrectStart time.Time, offset time.Duration) time.Time {
	return incorrectStart.Add(-offset)
}

// IsValidDetection - Check whether a detection is within an allowed time window after start.
func IsValidDetection(start time.Time, from, to time.Duration, detection time.Time) bool {
	elapsed := detection.Sub(start)
	return elapsed >= from && elapsed <= to
}

// FormatResult - Format leaderboard output.
func FormatResult(result, leader time.Duration) string {
	diff := result - leader
	if diff < 0 {
		diff = 0
	}
	return fmt.Sprintf("%s +%.3f", result.String(), diff.Seconds())
}
