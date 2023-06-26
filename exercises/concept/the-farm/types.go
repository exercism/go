package thefarm

// This file contains types used in the exercise and tests and should not be modified.

// FodderCalculator provides helper methods to determine the optimal
// amount of fodder to feed cows.
type FodderCalculator interface {
	FodderAmount(int) (float64, error)
	FatteningFactor() (float64, error)
}
