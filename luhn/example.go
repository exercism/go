package luhn

import "strings"

func Valid(id string) bool {

	if len(strings.TrimSpace(id)) == 1 {
		return false
	}

	d := make([]int, 0, len(id))

	for _, r := range id {
		if r == ' ' {
			continue
		}
		if r < '0' || r > '9' {
			return false
		}
		d = append(d, int(r-'0'))
	}

	return sum(d)%10 == 0
}

func sum(d []int) (s int) {
	for i, x := range d {
		j := len(d) - i
		if j%2 == 0 {
			x *= 2
			if x > 9 {
				x -= 9
			}
		}
		s += x
	}
	return
}
