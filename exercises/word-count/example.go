package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is a map of the frequency of occurrence keyed to the unique word.
type Frequency map[string]int

// WordCount returns the map of frequency of words based on the input phrase.
func WordCount(phrase string) Frequency {
	freq := Frequency{}
	for _, word := range strings.Fields(normalize(phrase)) {
		word = strings.Trim(word, "'")
		freq[word]++
	}
	return freq
}

func normalize(phrase string) string {
	phrase = strings.ToLower(phrase)
	// allow for apostrophes in words
	return regexp.MustCompile(`[^\w|']`).ReplaceAllLiteralString(phrase, " ")
}
