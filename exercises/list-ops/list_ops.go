package listops

// IntList is an abstraction of a list of integers
type IntList []int

// PredFunc function to filter values
type PredFunc func(int) bool

// BinFunc function to combine values
type BinFunc func(int, int) int

// UnaryFunc function to transform values
type UnaryFunc func(int) int

// Foldl folds (reduces) the given list from the left with a function
func (ls IntList) Foldl(fn BinFunc, initial int) int {
	panic("TODO: implement me")
}

// Foldr folds (reduces) the given list from the right with a function
func (ls IntList) Foldr(fn BinFunc, initial int) int {
	panic("TODO: implement me")
}

// Filter list returning only values that satisfy the filter function
func (ls IntList) Filter(fn PredFunc) IntList {
	panic("TODO: implement me")
}

// Length returns the length of a list
func (ls IntList) Length() int {
	panic("TODO: implement me")
}

// Map returns a list of elements whose values equal the list value transformed by the mapping function
func (ls IntList) Map(fn UnaryFunc) IntList {
	panic("TODO: implement me")
}

// Reverse reverses the list
func (ls IntList) Reverse() IntList {
	panic("TODO: implement me")
}

// Append adds the elements of the argument list to the receiver
func (ls IntList) Append(other IntList) IntList {
	panic("TODO: implement me")
}

// Concat concatenates a list of lists
func (ls IntList) Concat(others []IntList) IntList {
	panic("TODO: implement me")
}
