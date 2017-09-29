package say

import "errors"

const testVersion = 1

var small = []string{"zero", "one", "two", "three", "four", "five", "six",
	"seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen",
	"fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens = []string{"ones", "ten", "twenty", "thirty", "forty",
	"fifty", "sixty", "seventy", "eighty", "ninety"}
var scale = []string{"thousand", "million", "billion"}

func Say(n uint64) (string, error) {
	if n >= 1000000000000 {
		return "", errors.New("value out of range")
	}
	switch {
	case n < 20:
		return small[n], nil
	case n < 100:
		t := tens[n/10]
		s := n % 10
		if s > 0 {
			t += "-" + small[s]
		}
		return t, nil
	case n < 1000:
		h := small[n/100] + " hundred"
		s := n % 100
		if s > 0 {
			temp, _ := Say(s)
			h += " " + temp
		}
		return h, nil
	}
	sx := ""
	if p := n % 1000; p > 0 {
		sx, _ = Say(p)
	}
	for i := 0; n >= 1000; i++ {
		n /= 1000
		if p := n % 1000; p > 0 {
			temp, _ := Say(p)
			ix := temp + " " + scale[i]
			if sx > "" {
				ix += " " + sx
			}
			sx = ix
		}
	}
	return sx, nil
}
