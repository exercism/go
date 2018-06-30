package listops

type IntSlice []int
type predFunc func(int) bool
type binFunc func(int, int) int

// Foldl applies a left fold to the list given a binary function and an initial
// value
func (s IntSlice) Foldl(fn binFunc, initial int) int {
	if len(s) == 0 {
		return initial
	}
	x, xs := s[0], s[1:]
	return xs.Foldl(fn, fn(initial, x))
}

func (s IntSlice) Foldr(fn binFunc, initial int) int {
	flippedFunc := func(x, y int) int { return fn(y, x) }
	return reverseInt(s).Foldl(flippedFunc, initial)
}
func reverseInt(ints IntSlice) IntSlice {
	c := make([]int, len(ints))
	c = append([]int(nil), ints...)
	for left, right := 0, len(ints)-1; left < right; left, right = left+1, right-1 {
		c[left], c[right] = c[right], c[left]
	}
	return c
}

func (s IntSlice) Filter(fn predFunc) IntSlice {
	filtered := make([]int, 0, len(s))
	var filterAcc func(predFunc, IntSlice, IntSlice) IntSlice
	filterAcc = func(fn predFunc, acc IntSlice, lst IntSlice) IntSlice {
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

func (s IntSlice) Length() int {
	// anything else is just ridiculous
	return len(s)
}
