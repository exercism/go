package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	panic("Please implement the Foldl function")
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	panic("Please implement the Foldl function")
}

func (s IntList) Filter(fn func(int) bool) IntList {
	panic("Please implement the Filter function")
}

func (s IntList) Length() int {
	panic("Please implement the Length function")
}

func (s IntList) Map(fn func(int) int) IntList {
	panic("Please implement the Map function")
}

func (s IntList) Reverse() IntList {
	panic("Please implement the Reverse function")
}

func (s IntList) Append(lst IntList) IntList {
	panic("Please implement the Append function")
}

func (s IntList) Concat(lists []IntList) IntList {
	panic("Please implement the Concat function")
}
