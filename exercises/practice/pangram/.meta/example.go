package pangram

import "strings"

// IsPangram checks if given string is a pangram
func IsPangram(s string) bool {
	var check [26]bool
	var count int
	for _, c := range strings.ToLower(s) {
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
