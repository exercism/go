package sieve

func Sieve(limit int) (primes []int) {
	c := make([]bool, limit)
	for p := 2; p < limit; {
		for i := p + p; i < limit; i += p {
			c[i] = true
		}
		for p++; p < limit && c[p]; p++ {
		}
	}
	for i := 2; i < limit; i++ {
		if !c[i] {
			primes = append(primes, i)
		}
	}
	return
}
