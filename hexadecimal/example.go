package hexadecimal

import (
	"math"
	"strconv"
	"strings"
)

//  The easy way:
//
//  func ToDecimal(s string) int64 {
//    v, _ := strconv.ParseInt(s, 16, 64)
//    return v
//  }

//  A bit harder

func ToDecimal(s string) (n int64) {

	if invalid(s) {
		return int64(0)
	}

	length := len(s)

	for i := length; i > 0; i-- {
		v := intValue(s[i-1])
		n = n + (v * int64(math.Pow(16, float64(length-i))))
	}

	return n
}

func invalid(s string) bool {
	for _, char := range strings.Split(s, "") {
		if char > "f" {
			return true
		}
	}
	return false
}

func intValue(c uint8) int64 {
	v, _ := strconv.ParseInt(string(c), 16, 64)
	return v
}
