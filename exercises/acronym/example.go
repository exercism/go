package acronym

import (
	"regexp"
	"strings"
)

func Abbreviate(s string) string {
	regex := regexp.MustCompile("[A-Z]+[a-z]*|[a-z]+")
	words := regex.FindAllString(s, -1)

	var abbr []string
	for _, word := range words {
		abbr = append(abbr, string(word[0]))
	}

	return strings.ToUpper(strings.Join(abbr, ""))
}
