package leap

// IsLeapYear confirms whether a given year is a leap year.
func IsLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

// This test versioning is specific to Exercism,
// you don't need to worry about it now.
const testVersion = 3
