package thefarm

import "errors"

// DivideFood computes the fodder amount for the given cows
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodder, err := weightFodder()
	if errors.Is(err, ErrWeight) {
		fodder *= 2
	} else {
		if err != nil {
			return 0, err
		}
	}
	if fodder < 0 {
		return 0, errors.New("Negative fodder")
	}
	if cows == 0 {
		return 0, errors.New("Division by zero")
	}
	if cows < 0 {
		return 0, errors.New("Negative cows")
	}
	return fodder / float64(cows), nil
}
