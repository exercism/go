package strain

type Ints []int

func (s Ints) Keep(f func(int) bool) (r Ints) {
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return
}

func (s Ints) Discard(f func(int) bool) (r Ints) {
	for _, e := range s {
		if !f(e) {
			r = append(r, e)
		}
	}
	return
}

type Strings []string

func (s Strings) Keep(f func(string) bool) (r Strings) {
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return
}

type Lists [][]int

func (s Lists) Keep(f func([]int) bool) (r Lists) {
	for _, e := range s {
		if f(e) {
			r = append(r, e)
		}
	}
	return
}
