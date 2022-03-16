package cards

// AllItems returns a slice with the numbers 1..10 in ascending order
func AllItems() []int {
	return []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of a range we want it to return -1
func GetItem(slice []int, index int) int {
	if len(slice) <= index || index < 0 {
		return -1
	}
	return slice[index]
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if len(slice) <= index || index < 0 {
		return append(slice, value)
	}
	slice[index] = value
	return slice
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
	if len(slice) <= index || index < 0 {
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

// PrependItems adds an arbitrary number of values at the front of a slice
func PrependItems(slice []int, value ...int) []int {
	return append(value, slice...)
}
