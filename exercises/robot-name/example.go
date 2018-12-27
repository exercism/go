// +build !bonus

package robotname

import (
	"fmt"
	"math/rand"
)

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		r.name = fmt.Sprintf("%c%c%03d",
			'A'+byte(rand.Intn(26)),
			'A'+byte(rand.Intn(26)),
			rand.Intn(1000))
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	*r = Robot{}
}
