package thefarm

import "errors"

// WeightFodder returns the amount of available fodder.
type WeightFodder func() (float64, error)

// ErrWeight
var ErrWeight = errors.New("sensor error")

// DivideFood computes the fodder amount for the given cows
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	panic("Please implement DivideFood(float64, int) (float64, error)")
}
