package ocr

import "strings"

var digit = map[string]byte{
	" _ | ||_|   ": '0',
	"     |  |   ": '1',
	" _  _||_    ": '2',
	" _  _| _|   ": '3',
	"   |_|  |   ": '4',
	" _ |_  _|   ": '5',
	" _ |_ |_|   ": '6',
	" _   |  |   ": '7',
	" _ |_||_|   ": '8',
	" _ |_| _|   ": '9',
}

func recognizeDigit(lines []string, start int) byte {
	if len(lines) < 4 {
		return '?'
	}
	need := start + 3
	for _, l := range lines {
		if len(l) < need {
			return '?'
		}
	}
	if d, ok := digit[lines[0][start:need]+
		lines[1][start:need]+
		lines[2][start:need]+
		lines[3][start:need]]; ok {
		return d
	}
	return '?'
}

func Recognize(s string) []string {
	lines := strings.Split(s, "\n")
	if len(lines[0]) == 0 {
		lines = lines[1:]
	}
	r := make([]string, (len(lines)+3)/4)
	for i := range r {
		max := 0
		for l := 0; l < 4 && l < len(lines); l++ {
			if n := len(lines[l]); n > max {
				max = n
			}
		}
		b := make([]byte, (max+2)/3)
		for j := range b {
			b[j] = recognizeDigit(lines, j*3)
		}
		r[i] = string(b)
		lines = lines[4:]
	}
	return r
}
