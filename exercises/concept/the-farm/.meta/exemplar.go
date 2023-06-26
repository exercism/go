package thefarm

import (
	"errors"
	"fmt"
)

// DivideFood uses FodderCalculator methods to determine the amount of fodder each cow
// should get.
func DivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
	fodder, err := fodderCalculator.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	factor, err := fodderCalculator.FatteningFactor()
	if err != nil {
		return 0, err
	}

	return fodder / float64(cows) * factor, nil
}

// ValidateInputAndDivideFood does the same as DivideFood but it checks that
// the number of cows is positive first and returns an error if not.
func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(fodderCalculator, cows)
}

// InvalidCowsError is a custom error type that can be used
// to signal invalid input.
type InvalidCowsError struct {
	cows    int
	message string
}

// Error implements the error interface.
func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

// ValidateNumberOfCows returns an InvalidCowsError when the number of
// cows passed in was not positive and nil otherwise.
func ValidateNumberOfCows(cows int) error {
	if cows < 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "there are no negative cows",
		}
	}

	if cows == 0 {
		return &InvalidCowsError{
			cows:    cows,
			message: "no cows don't need food",
		}
	}

	return nil
}
