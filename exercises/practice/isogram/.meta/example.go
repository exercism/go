package isogram

import "strings"

func IsIsogram(word string) bool {
	lowerWord := strings.ToLower(word)
	for i, r := range lowerWord {
		if r != ' ' && r != '-' && i < strings.LastIndex(lowerWord, string(r)) {
			return false
		}
	}
	return true
}
