package summultiples

// MultipleSummer returns the sum of the multiples of the given divisors
// up to, but not including, the given limit.
func MultipleSummer(limit int, divisors ...int) (sum int) {
	for i := 1; i < limit; i++ {
		for _, d := range divisors {
			if i%d == 0 {
				sum += i
				break
			}
		}
	}
	return
}
