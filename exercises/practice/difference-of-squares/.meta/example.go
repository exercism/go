package diffsquares

func SquareOfSum(n int) int {
	s := 0
	for i := 1; i <= n; i++ {
		s += i
	}
	return s * s
}

func SumOfSquares(n int) (s int) {
	for i := 1; i <= n; i++ {
		s += i * i
	}
	return
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
