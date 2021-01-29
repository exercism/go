package wordy

import (
	"strconv"
	"strings"
)

func Answer(q string) (a int, ok bool) {
	w := strings.Fields(q)
	if len(w) < 3 { // length check for first two words and last word
		return
	}
	if w[0] != "What" || w[1] != "is" { // first two words required
		return
	}
	w = w[2:]
	last := len(w) - 1
	wl := w[last]
	if wl[len(wl)-1] != '?' { // trailing ? required
		return
	}
	w[last] = wl[:len(wl)-1]
	a, err := strconv.Atoi(w[0]) // first word after preamble must be a number
	if err != nil {
		return
	}
	for i := 1; i < len(w); i++ { // remainder of q must alternate between op, number
		op := w[i]
		i++
		switch op {
		case "multiplied", "divided":
			if i == len(w) || w[i] != "by" { // length check, required word
				return
			}
			i++ // consume
		}
		if i == len(w) { // length check for operand
			return
		}
		x, err := strconv.Atoi(w[i]) // must be a number
		if err != nil {
			return
		}
		switch op { // apply operator
		case "plus":
			a += x
		case "minus":
			a -= x
		case "multiplied":
			a *= x
		case "divided":
			a /= x
		default: // valid operator not found
			return
		}
	}
	return a, true
}
