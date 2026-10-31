package genericset

// Set is implemented using a map.
type Set[T comparable] map[T]bool

// NewSet instantiates a Set.
func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

// Equal checks if two sets are the same.
// Two sets are the same if they each have all the elements in the other.
func (s Set[T]) Equal(other Set[T]) bool {
	for i, _ := range s {
		if !other.Has(i) {
			return false
		}
	}
	for i := range other {
		if !s.Has(i) {
			return false
		}
	}
	return true
}

// Add an element to a Set.
func (s Set[T]) Add(elem T) {
	s[elem] = true
}

// Check if an element is in a Set.
func (s Set[T]) Has(elem T) bool {
	return s[elem]
}

// Remove an element from a Set.
func (s Set[T]) Remove(elem T) {
	delete(s, elem)
}

// Check if a Set is empty.
func (s Set[T]) IsEmpty() bool {
	return len(s) == 0
}
