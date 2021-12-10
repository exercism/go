package cards

// GetItem retrieves an item from a slice at given position. The second return value indicates whether
// the given index exists in the slice or not.
func GetItem(slice []int, index int) (int, bool) {
	panic("Please implement the GetItem function")
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	panic("Please implement the SetItem function")
}

// PrefilledSlice creates a slice of given length and prefills it with the given value.
func PrefilledSlice(value, length int) []int {
	panic("Please implement the PrefilledSlice function")
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	panic("Please implement the RemoveItem function")
}
