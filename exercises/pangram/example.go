package pangram

import (
	"strings"
)


func IsPangram(s string) bool {
	lowerString := strings.ToLower(s)
	var check [26]bool
	var count int
	for _, c := range lowerString {
		if c >= 'a' {
			if c > 'z' {
				continue
			}
			c -= 97
		} else {
			continue
		}

		if !check[c] {
			if count == 25 {
				return true
			}
			check[c] = true
			count++
		}
	}
	return false
}
