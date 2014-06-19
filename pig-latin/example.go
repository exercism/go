package igpay

import (
	"regexp"
	"strings"
)

var vowel = regexp.MustCompile(`^([aeiou]|y[^aeiou]|xr)[a-z]*`)
var cons = regexp.MustCompile(`^([^aeiou]?qu|[^aeiou]+)([a-z]*)`)

func PigLatin(s string) string {
	sw := strings.Fields(s)
	for i, w := range sw {
		l := strings.ToLower(w)
		if vowel.MatchString(l) {
			sw[i] = l + "ay"
		} else if x := cons.FindStringSubmatchIndex(l); x != nil {
			sw[i] = l[x[3]:] + l[:x[3]] + "ay"
		}
	}
	return strings.Join(sw, " ")
}
