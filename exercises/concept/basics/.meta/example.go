package basics

// OvenTime returns the expected baking time in minutes.
func OvenTime() int {
	return 40
}

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actual int) int {
	return OvenTime() - actual
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

// TotalTime calculates the total time needed to create and bake a lasagna.
func TotalTime(numberOfLayers, actualMinutesInOven int) int {
	return PreparationTime(numberOfLayers) + actualMinutesInOven
}
