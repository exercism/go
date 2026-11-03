package genericset

// Implement Set[T comparable] as a map[T]struct{}.
type Set[T comparable] map[T]struct{}

func NewSet[T comparable]() Set[T] {
	panic("Please implement the NewSet function")
}

// Add an element to a Set.
func (s Set[T]) Add(elem T) {
	panic("Please implement the Add function")
}

// Check if an element is in a Set.
func (s Set[T]) Contains(elem T) bool {
	panic("Please implement the Contains function")
}

// Check if two sets have the same elements.
func (s Set[T]) Equals(other Set[T]) bool {
	panic("Please implement the Equals function")
}
