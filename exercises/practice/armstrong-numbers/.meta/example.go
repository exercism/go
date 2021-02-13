package armstrong

import "math"

func order(n int) (result int) {
	for n != 0 {
		result++
		n /= 10
	}
	return
}

// IsNumber returns true for an armstrong number
func IsNumber(n int) bool {
	originalNumber := n
	pow := order(n)
	sum := 0
	for n != 0 {
		remainder := n % 10
		sum += int(math.Pow(float64(remainder), float64(pow)))
		n /= 10
	}
	return originalNumber == sum
}
