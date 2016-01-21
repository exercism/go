package leap

const TestVersion = 1

func IsLeapYear(i int) bool {
	return i%4 == 0 && i%100 != 0 || i%400 == 0
}
