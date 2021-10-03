package thefarm

import "errors"

// DivideFood computes the fodder amount per cow for the given cows
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder.FodderAmount()
	if err != nil {
		if !errors.Is(err, ScaleError{}) {
			return 0, err
		}
		fodder *= 2
	}
	if fodder < 0 {
		return 0, errors.New("Negative fodder")
	}
	if cows == 0 {
		return 0, errors.New("Division by zero")
	}
	if cows < 0 {
		return 0, SillyNephewError{cows}
	}
	return fodder / float64(cows), nil
}
