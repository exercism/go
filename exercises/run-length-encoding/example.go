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
				count, output = encode(count, i, s, output)
			}
		}
	}
	if len(s) != 0 {
		_, output = encode(count, len(s), s, output)
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

func encode(count, i int, s, output string) (int, string) {
	if count > 1 {
		output += fmt.Sprintf("%d%c", count, s[i-1])
		return 1, output
	}
	output += fmt.Sprintf("%c", s[i-1])
	return 1, output
}
