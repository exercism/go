package zero

// ZeroBool returns a bool set to its zero value.
func ZeroBool() bool {
	return false
}

// ZeroInt returns an int set to its zero value.
func ZeroInt() int {
	return 0
}

// ZeroString returns a string set to its zero value.
func ZeroString() string {
	return ""
}

// ZeroFunc returns a func set to its zero value.
func ZeroFunc() func() {
	return nil
}

// ZeroPointer returns a pointer set to its zero value.
func ZeroPointer() *int {
	return nil
}

// ZeroMap returns a map set to its zero value.
func ZeroMap() map[int]int {
	return nil
}

// ZeroSlice returns a slice set to its zero value.
func ZeroSlice() []int {
	return nil
}

// ZeroChannel returns a channel set to its zero value.
func ZeroChannel() chan int {
	return nil
}

// ZeroInterface returns an interface set to its zero value.
func ZeroInterface() interface{} {
	return nil
}
