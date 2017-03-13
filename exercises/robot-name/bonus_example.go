// +build bonus

package robotname

import (
	"fmt"
	"math/rand"
)

type Robot struct {
	name string
}

var issued = map[string]bool{}

func (r *Robot) Name() string {
	if r.name > "" {
		return r.name
	}
	a1 := rand.Intn(26)
	a2 := rand.Intn(26)
	n := rand.Intn(1000)
	start := fmt.Sprintf("%c%c%03d", 'A'+byte(a1), 'A'+byte(a2), n)
	r.name = start
	for issued[r.name] {
		if n = (n + 1) % 1000; n == 0 {
			if a2 = (a2 + 1) % 26; a2 > 0 {
				a1 = (a1 + 1) % 26
			}
		}
		r.name = fmt.Sprintf("%c%c%03d", 'A'+byte(a1), 'A'+byte(a2), n)
		if r.name == start {
			panic("all valid robot names issued")
		}
	}
	issued[r.name] = true
	return r.name
}

func (r *Robot) Reset() {
	*r = Robot{}
}
