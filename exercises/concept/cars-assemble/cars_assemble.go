package cars

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	panic("CalculateProductionRatePerHour not implemented")
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	panic("CalculateProductionRatePerMinute not implemented")
}

// successRate is used to calculate the ratio of an item being created without
// error for a given speed
func successRate(speed int) float64 {
	if speed == 0 {
		return 0.0
	}

	if speed < 5 {
		return 1.0
	}

	if speed >= 9 {
		return 0.77
	}

	return 0.9
}
