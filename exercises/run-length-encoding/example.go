package encode

import (
	"fmt"
	"strconv"
)

func RunLengthEncode(s string) string {
	count := 1
	var output string
	for i, c := range s {
		if i != 0 {
			if rune(s[i-1]) == c {
				count++
			} else {
				count, output = 1, output+encode(count, i, s)
			}
		}
	}
	if s != "" {
		_, output = 1, output+encode(count, len(s), s)
	}
	return output
}

func RunLengthDecode(s string) string {
	count := 1
	var stringCount, output string
	for _, c := range s {
		if _, err := strconv.Atoi(string(c)); err == nil {
			stringCount += string(c)
		} else {
			if stringCount != "" {
				count, _ = strconv.Atoi(stringCount)
			}
			for j := 0; j < count; j++ {
				output += string(c)
			}
			count = 1
			stringCount = ""
		}
	}
	return output
}

func encode(count, i int, s string) string {
	if count > 1 {
		return fmt.Sprintf("%d%c", count, s[i-1])
	}
	return fmt.Sprintf("%c", s[i-1])
}
