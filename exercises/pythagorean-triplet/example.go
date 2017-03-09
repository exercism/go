package pythagorean

const testVersion = 1

type Triplet [3]int

func pyth(a, b, c int) bool {
	return a*a+b*b == c*c
}

func Range(min, max int) (ts []Triplet) {
	for a := min; a <= max; a++ {
		for b := a; b <= max; b++ {
			for c := b; c <= max; c++ {
				if pyth(a, b, c) {
					ts = append(ts, Triplet{a, b, c})
				}
			}
		}
	}
	return
}

func Sum(sum int) (ts []Triplet) {
	max := sum / 2
	for a := 1; a <= max; a++ {
		for b := a; b <= max; b++ {
			if c := sum - a - b; pyth(a, b, c) {
				ts = append(ts, Triplet{a, b, c})
			}
		}
	}
	return
}
