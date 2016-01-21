package summultiples

func MultipleSummer(divisors ...int) func(int) int {
	return func(limit int) int {
		sum := 0
		for i := 1; i < limit; i++ {
			for _, d := range divisors {
				if i%d == 0 {
					sum += i
					break
				}
			}
		}
		return sum
	}
}
