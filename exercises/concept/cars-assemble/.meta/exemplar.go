package cars

// CalculateWorkingCarsPerHour for the assembly line, taking into account
// the production rate (in units per hour) and the success rate (as a percentage)
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return float64(productionRate) * successRate / 100
}

// CalculateWorkingCarsPerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	return int(CalculateWorkingCarsPerHour(productionRate, successRate) / 60)
}

// CalculateCost works out how much the materials to produce a
// number of cars cost
func CalculateCost(carsCount int) uint {
	tens := carsCount / 10
	afterTens := carsCount % 10
	threes := afterTens / 3
	singles := afterTens % 3
	return uint(tens*95000 + threes*29000 + singles*10000)
}
