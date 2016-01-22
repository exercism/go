package wordcount

import (
	"regexp"
	"strings"
)

const TestVersion = 1

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := Frequency{}
	for _, word := range strings.Fields(normalize(phrase)) {
		freq[word]++
	}
	return freq
}

func normalize(phrase string) string {
	r, _ := regexp.Compile(`[^\w]`)
	phrase = strings.ToLower(phrase)
	return r.ReplaceAllLiteralString(phrase, " ")
}
