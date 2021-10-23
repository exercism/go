package isbn

import (
	"errors"
	"math"
	"strconv"
	"unicode"
)

func IsValidISBN(isbn string) bool {

	isbn = dropHyphen(isbn)

	ary, err := strToSlice(isbn)
	if len(ary) != 10 || err != nil {
		return false
	}

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

	for pos, char := range isbn {
		switch {
		case unicode.IsLetter(char) && (char != 'X' || pos != 9):
			err = errors.New("invalid character")
			return
		case char == 'X':
			result = append(result, 10)
		default:
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
