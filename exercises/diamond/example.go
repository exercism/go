package diamond

import (
	"errors"
	"strings"
)

const testVersion = 1

const startIndex = 'A'

// Gen builds a diamond
func Gen(char byte) (string, error) {
	if char > 'Z' || char < 'A' {
		return "", errors.New(string(char) + " isn't supported, only between (A, Z)")
	}
	return gen(char), nil
}

func gen(char byte) string {
	var output []string
	currentIndex := int(char - startIndex)
	for i := 0; i <= currentIndex; i++ {
		output = append(output, getLine(currentIndex, i))
	}
	for i := currentIndex - 1; i > -1; i-- {
		output = append(output, getLine(currentIndex, i))
	}
	return strings.Join(output, "\n") + "\n"
}

func getLine(currentStart, current int) string {
	diff := currentStart - current
	return spaces(diff) + alphabets(current) + spaces(diff)
}

func alphabets(current int) string {
	if current == 0 {
		return "A"
	}
	c := current + startIndex
	return string(c) + spaces(current*2-1) + string(c)
}

func spaces(n int) string {
	return strings.Repeat(" ", n)
}
