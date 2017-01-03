package wordcount

import (
	"regexp"
	"strings"
)

const testVersion = 3

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
	//  Allow for apostrophes in words.
	r, _ := regexp.Compile(`[^\w|']`)
	phrase = strings.ToLower(phrase)
	return r.ReplaceAllLiteralString(phrase, " ")
}
