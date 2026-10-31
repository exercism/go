package genericset

// Implement Set[T comparable] as a map[T]bool or map[T]interface{}.
//
// Define the Set type here.

func NewSet[T comparable]() Set[T] {
	panic("Please implement the NewSet function")
}

// Equal checks if two sets are the same.
// Two sets are the same if they each have all the elements in the other.
func (s Set[T]) Equal(other Set[T]) bool {
	panic("Please implement the Equal function")
}

// Add an element to a Set.
func (s Set[T]) Add(elem T) {
	panic("Please implement the Add function")
}

// Check if an element is in a Set.
func (s Set[T]) Has(elem T) bool {
	panic("Please implement the Has function")
}

// Remove an element from a Set.
func (s Set[T]) Remove(elem T) {
	panic("Please implement the Remove function")
}

// Check if a Set is empty.
func (s Set[T]) IsEmpty() bool {
	panic("Please implement the IsEmpty function")
}
