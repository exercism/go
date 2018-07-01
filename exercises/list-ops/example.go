package listops

//
type IntList []int
type predFunc func(int) bool
type binFunc func(int, int) int
type unaryFunc func(int) int

// Foldl folds (reduces) the given list from the left with a function
func (s IntList) Foldl(fn binFunc, initial int) int {
	if len(s) == 0 {
		return initial
	}
	x, xs := s[0], s[1:]
	return xs.Foldl(fn, fn(initial, x))
}

// Foldr folds (reduces) the given list from the right with a function
func (s IntList) Foldr(fn binFunc, initial int) int {
	flippedFunc := func(x, y int) int { return fn(y, x) }
	return reverseInt(s).Foldl(flippedFunc, initial)
}

func reverseInt(ints IntList) IntList {
	c := make([]int, len(ints))
	c = append([]int(nil), ints...)
	for left, right := 0, len(ints)-1; left < right; left, right = left+1, right-1 {
		c[left], c[right] = c[right], c[left]
	}
	return c
}

// Filter list returning only values that satisfy the filter function
func (s IntList) Filter(fn predFunc) IntList {
	filtered := make([]int, 0, len(s))
	var filterAcc func(predFunc, IntList, IntList) IntList
	filterAcc = func(fn predFunc, acc IntList, lst IntList) IntList {
		if len(lst) == 0 {
			return acc
		}
		x, xs := lst[0], lst[1:]
		if fn(x) {
			acc = append(acc, x)
		}
		return filterAcc(fn, acc, xs)
	}
	return filterAcc(fn, filtered, s)
}

// Length returns the length of a list
func (s IntList) Length() int {
	// anything else is just ridiculous
	return len(s)
}

// Map returns a list of elements whose values equal the list value transformed by the mapping function
func (s IntList) Map(fn unaryFunc) IntList {
	newSlice := make([]int, len(s))
	for idx, elt := range s {
		newSlice[idx] = fn(elt)
	}
	return newSlice
}

// Reverse reverses the list
func (s IntList) Reverse() IntList {
	last := len(s) - 1
	newSlice := make([]int, last+1)
	for idx, elt := range s {
		newSlice[last-idx] = elt
	}
	return newSlice
}

// Append adds the elements of the argument list to the receiver
func (s IntList) Append(lst IntList) IntList {
	offset := len(s)
	newSlice := make([]int, offset+len(lst))
	copy(newSlice, s)
	for idx, elt := range lst {
		newSlice[offset+idx] = elt
	}
	return newSlice
}

// Concat concatenates a list of lists
func (s IntList) Concat(lists []IntList) IntList {
	// totalLength := foldl(func(acc int, x IntList) int { return acc + x.Length() }, 0, lists)
	newSlice := make([]int, len(s))
	copy(newSlice, s)
	for _, l := range lists {
		newSlice = append(newSlice, l...)
	}
	return newSlice
}
