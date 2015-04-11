package prime

const TestVersion = 1

func Factors(n int64) []int64 {
	factors := []int64{}
	possibleFactor := int64(2)

	for possibleFactor*possibleFactor <= n {
		for n%possibleFactor == 0 {
			factors = append(factors, possibleFactor)
			n /= possibleFactor
		}
		possibleFactor++
	}

	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}
