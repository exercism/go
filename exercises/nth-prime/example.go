package prime

func Nth(n int) (p int, ok bool) {
	switch {
	case n < 1:
		return 0, false
	case n == 1:
		return 2, true
	}
	n--
	p = 3
	inc := 1
	sqr := 1
	sqrt := 1
	for {
		for f := 3; ; f += 2 {
			if f > sqrt {
				n--
				if n == 0 {
					return p, true
				}
				break
			}
			if p%f == 0 {
				break
			}
		}
		p += 2
		if p > sqr {
			inc += 2
			sqr += inc
			sqrt++
		}
	}
}
