package acronym

import (
	"fmt"
	"regexp"
	"strings"
)

const testVersion = 2

func Abbreviate(s string) string {
	regex := regexp.MustCompile("[A-Z]+[a-z]*|[a-z]+")
	words := regex.FindAllString(s, -1)
	abbr := []string{}

	for _, word := range words {
		abbr = append(abbr, string(word[0]))
	}

	return fmt.Sprintf("%s", strings.ToUpper(strings.Join(abbr, "")))
}
