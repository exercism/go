package diffsquares

func SquareOfSums(n int) int {
	s := 0
	for i := 0; i <= n; i++ {
		s += i
	}
	return s * s
}

func SumOfSquares(n int) (s int) {
	for i := 0; i <= n; i++ {
		s += i * i
	}
	return
}

func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
