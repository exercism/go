package perfect

import "errors"

// Classification is the category devised by Greek Mathematician Nicomachus
type Classification string

// Defined categories
const (
	ClassificationDeficient Classification = "deficient"
	ClassificationPerfect   Classification = "perfect"
	ClassificationAbundant  Classification = "abundant"
)

// Errors
var (
	ErrOnlyPositive = errors.New("Positive number is required")
)

// Classify finds the category of given natural number
func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return "", ErrOnlyPositive
	}
	var sum int64
	for i := int64(1); i < n; i++ {
		if n%i == 0 {
			if sum = sum + i; sum > n {
				return ClassificationAbundant, nil
			}
		}
	}
	if sum == n {
		return ClassificationPerfect, nil
	} else if sum < n {
		return ClassificationDeficient, nil
	}
	return ClassificationAbundant, nil
}
