package deep_thought

import (
	"errors"
	"fmt"
)

const answer = 42

// ErrCalculation is a general error that can be exported
var ErrCalculation = errors.New("calculation error")

// ComputerState determines if the deepthought computer is ready for its computation
func ComputerState(isReady bool) error {
	if !isReady {
		return errors.New("computer still booting")
	}
	return nil
}

// AnswerToLife checks if the given number is the answer to life
func AnswerToLife(number int) (bool, error) {
	if number != answer {
		return false, fmt.Errorf("%s: wrong answer to life: %s", ErrCalculation.Error(), number)
	}
	return true, nil
}

// AskComputer checks if number is the answer to life, depending on the state of the computer
func AskComputer(isReady bool, number int) (bool, error) {
	err := ComputerState(isReady)
	if err != nil {
		return false, fmt.Errorf("proposal %d could not be processed: %s", number, err)
	}
	status, err := AnswerToLife(number)
	return status, err
}
