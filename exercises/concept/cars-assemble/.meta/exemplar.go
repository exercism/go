package cars

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
func SuccessRate(speed int) float64 {
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

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	const defaultRate = 221.0

	rateForSpeed := defaultRate * float64(speed)

	return rateForSpeed * SuccessRate(speed)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60)
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	if rate := CalculateProductionRatePerHour(speed); rate <= limit {
		return rate
	}
	return limit
}
