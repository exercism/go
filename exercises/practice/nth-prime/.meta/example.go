package prime

import "errors"

var primes []int = []int{2}

func Nth(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("input must be greater than 1")
	}
	for i := 3; len(primes) < n; i++ {
		isPrime(i)
	}
	return primes[n-1], nil
}

func isPrime(n int) {
	for _, v := range primes {
		if n%v == 0 {
			return
		}
	}
	primes = append(primes, n)
}
