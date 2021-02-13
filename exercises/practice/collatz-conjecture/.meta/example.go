package collatzconjecture

import (
	"errors"
)

// CollatzConjecture is an example implementation of the collatz conjecture exercise.
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return -1, errors.New("Only positive numbers are allowed")
	}

	steps := 0
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
		steps++
	}

	return steps, nil
}
