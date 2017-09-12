package isogram

import "strings"

func IsIsogram(word string) bool {
	for i, r := range strings.ToLower(word) {
		if r != ' ' && r != '-' && i < strings.LastIndex(word, string(r)) {
			return false
		}
	}

	return true
}
