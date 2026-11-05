package genericset

type Set[T comparable] map[T]struct{}

// NewSet returns a new generic set.
func NewSet[T comparable]() Set[T] {
	return make(Set[T])
}

// Add an element to a Set.
func (s Set[T]) Add(elem T) {
	s[elem] = struct{}{}
}

// Check if an element is in a Set.
func (s Set[T]) Contains(elem T) bool {
	_, ok := s[elem]
	return ok
}

// Check if two sets have the same elements.
func (s Set[T]) Equals(other Set[T]) bool {
	for i := range s {
		if !other.Contains(i) {
			return false
		}
	}
	for i := range other {
		if !s.Contains(i) {
			return false
		}
	}
	return true
}
