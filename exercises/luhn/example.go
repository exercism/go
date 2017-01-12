package luhn

const testVersion = 1

func Valid(n string) bool {
	if len(n) == 1 {
		return false
	}
	d := extract(n)
	if len(d) == 0 {
		return false
	}
	last := len(d) - 1
	return (check(d[:last])+d[last])%10 == 0
}

func extract(n string) []int {
	d := make([]int, 0, len(n))
	for _, r := range n {
		if r >= '0' && r <= '9' {
			d = append(d, int(r-'0'))
		}
	}
	return d
}

func check(d []int) int {
	for i := len(d) - 1; i >= 0; i -= 2 {
		x := 2 * d[i]
		if x > 9 {
			x -= 9
		}
		d[i] = x
	}
	s := 0
	for _, x := range d {
		s += x
	}
	return s
}
