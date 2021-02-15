package accumulate

import (
	"strings"
)

func Accumulate(s []string, f func(st string) string) (result []string) {
	return
}

func echo(c string) string {
	return c
}

func capitalize(word string) string {
	return strings.Title(word)
}
