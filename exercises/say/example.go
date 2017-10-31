package say

var small = []string{"zero", "one", "two", "three", "four", "five", "six",
	"seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen",
	"fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}
var tens = []string{"ones", "ten", "twenty", "thirty", "forty",
	"fifty", "sixty", "seventy", "eighty", "ninety"}
var scale = []string{"thousand", "million", "billion",
	"trillion", "quadrillion", "quintillion"}

func Say(n int64) (value string, ok bool) {
	if n < 0 || n > 999999999999 {
		return
	}
	return say(n), true
}

func say(n int64) string {
	switch {
	case n < 20:
		return small[n]
	case n < 100:
		t := tens[n/10]
		s := n % 10
		if s > 0 {
			t += "-" + small[s]
		}
		return t
	case n < 1000:
		h := small[n/100] + " hundred"
		s := n % 100
		if s > 0 {
			h += " " + say(s)
		}
		return h
	}
	sx := ""

	if p := n % 1000; p > 0 {
		sx = say(p)
	}
	for i := 0; n >= 1000; i++ {
		n /= 1000
		if p := n % 1000; p > 0 {
			ix := say(p) + " " + scale[i]
			if sx > "" {
				ix += " " + sx
			}
			sx = ix
		}
	}
	return sx
}
