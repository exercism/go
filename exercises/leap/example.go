package leap

const testVersion = 2

func IsLeapYear(i int) bool {
	return i%4 == 0 && i%100 != 0 || i%400 == 0
}
