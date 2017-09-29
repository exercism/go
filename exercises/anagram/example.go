package anagram

import (
	"sort"
	"strings"
)

func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)
	var matches []string
	for _, c := range candidates {
		if tc := strings.ToLower(c); isAnagram(subject, tc) {
			matches = append(matches, c)
		}
	}
	return matches
}

func isAnagram(subject, candidate string) bool {
	return subject != candidate && alphagram(subject) == alphagram(candidate)
}

func alphagram(s string) string {
	chars := strings.Split(s, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}
