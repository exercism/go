package isbn

import (
	"fmt"
	"math"
	"strconv"
)

func IsVaildISBN(isbn string) bool {

	// step 1
	isbn = dropHyphen(isbn)
	fmt.Println("step1: ", isbn)

	// step 2
	ary := strToAry(isbn)
	fmt.Println("step2: ", ary)

	if len(ary) != 10 {
		return false
	}

	// step 3
	return calcCheckDigit(ary)

}

func dropHyphen(isbn string) string {
	var result string
	for _, char := range isbn {
		if char == '-' {
			continue
		}
		result += string(char)
	}
	return result
}

func strToAry(isbn string) []int {
	var result []int
	for _, char := range isbn {
		if char == 'X' {
			result = append(result, 10)
		} else {
			s := string(char)
			i, _ := strconv.Atoi(s)
			result = append(result, i)
		}
	}

	return result
}

func calcCheckDigit(isbn []int) bool {
	var pool int
	for idx, value := range isbn {
		pool += int(math.Abs(float64(idx)-10)) * value
	}
	result := pool % 11

	return result == 0
}
