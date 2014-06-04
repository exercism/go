package prime

func Factors(n int) []int {
	factors := []int{}
	possibleFactor := 2

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
