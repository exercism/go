package thefarm

import (
	"errors"
	"fmt"
)

// SillyNephewError should be returned if your nephew thinks there are negative cows.
type SillyNephewError struct {
	Cows int
}

// Error implements the error interface.
func (sn *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", sn.Cows)
}

// DivideFood computes the fodder amount per cow for the given cows
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	if err != nil && err != ErrScaleMalfunction {
		return 0, err
	}
	if err == ErrScaleMalfunction {
		fodder *= 2
	}
	if fodder < 0 {
		return 0, errors.New("Negative fodder")
	}
	if cows == 0 {
		return 0, errors.New("Division by zero")
	}
	if cows < 0 {
		return 0, &SillyNephewError{Cows: cows}
	}

	return fodder / float64(cows), nil
}
