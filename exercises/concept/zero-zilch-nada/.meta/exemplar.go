package zero

// IsZeroBool determines if a given bool is set to its zero value.
func IsZeroBool(b bool) bool {
	if b == false {
		return true
	}
	return false
}

// IsZeroInt determines if a given int is set to its zero value.
func IsZeroInt(i int) bool {
	if i == 0 {
		return true
	}
	return false
}

// IsZeroString determines if a given string is set to its zero value.
func IsZeroString(s string) bool {
	if s == "" {
		return true
	}
	return false
}

// IsZeroFunc determines if a given function is set to its zero value.
func IsZeroFunc(f func()) bool {
	if f == nil {
		return true
	}
	return false
}

// IsZeroPointer determines if a given pointer is set to its zero value.
func IsZeroPointer(i *int) bool {
	if i == nil {
		return true
	}
	return false
}

// IsZeroMap determines if a given map is set to its zero value.
func IsZeroMap(m map[string]int) bool {
	if m == nil {
		return true
	}
	return false
}

// IsZeroSlice determines if a given slice is set to its zero value.
func IsZeroSlice(i []int) bool {
	if i == nil {
		return true
	}
	return false
}

// IsZeroChannel determines if a given channel is set to its zero value.
func IsZeroChannel(c chan int) bool {
	if c == nil {
		return true
	}
	return false
}

// IsZeroInterface determines if a given interface is set to its zero value.
func IsZeroInterface(i interface{}) bool {
	if i == nil {
		return true
	}
	return false
}
