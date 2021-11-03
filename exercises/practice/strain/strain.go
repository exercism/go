package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	panic("Please implement the Keep function")
}

func (i Ints) Discard(filter func(int) bool) Ints {
	panic("Please implement the Discard function")
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	panic("Please implement the Keep function")
}

func (s Strings) Keep(filter func(string) bool) Strings {
	panic("Please implement the Keep function")
}
