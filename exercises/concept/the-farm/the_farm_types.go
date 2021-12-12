package thefarm

import "fmt"

// This file contains types used in the exercise but should not be modified.

// WeightFodder returns the amount of available fodder.
type WeightFodder interface {
	FodderAmount() (float64, error)
}

// ScaleError indicates an error with the scale
type ScaleError struct{}

func (se ScaleError) Error() string {
	return "sensor error"
}

// SillyNephewError should be returned if your nephew thinks there are negative cows
type SillyNephewError struct {
	Cows int
}

func (sn SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", sn.Cows)
}
