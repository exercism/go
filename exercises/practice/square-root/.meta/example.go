package squareroot

import "errors"

func SquareRoot(number int) (int, error) {
	for i := range number + 1 {
		if i*i == number {
			return i, nil
		}
	}
	return 0, errors.New("no squareroot found")
}
