package atbash

import (
	"regexp"
	"strings"
)

var alphabet = "abcdefghijklmnopqrstuvwxyz"

func Atbash(s string) string {
	return chunk(convert(normalize(s)))
}

func chunk(s string) string {
	value := regexp.MustCompile(".{1,5}").FindAllString(s, -1)
	return strings.Join(value, " ")
}

func convert(s string) string {
	key := reverse(alphabet)
	inputSlice := strings.Split(s, "")
	originalSlice := strings.Split(alphabet, "")
	reversedSlice := strings.Split(key, "")
	result := ""
	for i := 0; i < len(s); i++ {
		char := inputSlice[i]
		index := indexOf(originalSlice, char)
		if index > -1 {
			result += reversedSlice[index]
		} else {
			result += char
		}
	}
	return result
}

func normalize(s string) string {
	s = strings.ToLower(s)
	return regexp.MustCompile("[^a-z0-9]").ReplaceAllString(s, "")
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func indexOf(slice []string, s string) int {
	for p, v := range slice {
		if v == s {
			return p
		}
	}
	return -1
}
