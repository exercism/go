package isbn

import (
	"errors"
	"math"
	"strconv"
	"unicode"
)

func IsValidISBN(isbn string) bool {

	// step 1
	isbn = dropHyphen(isbn)

	// step 2
	ary, err := strToSlice(isbn)

	if len(ary) != 10 || err != nil {
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

func strToSlice(isbn string) (result []int, err error) {

	for _, char := range isbn {
		if unicode.IsLetter(char) && char != 'X' {
			err = errors.New("invalid character")
			return
		} else if char == 'X' {
			result = append(result, 10)
		} else {
			i, _ := strconv.Atoi(string(char))
			result = append(result, i)
		}
	}
	return
}

func calcCheckDigit(isbn []int) bool {
	var pool int
	for idx, value := range isbn {
		pool += int(math.Abs(float64(idx)-10)) * value
	}
	result := pool % 11

	return result == 0
}
