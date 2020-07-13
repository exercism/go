package _meta

// EmptyInterface returns an empty (zero value) interface
func EmptyInterface() interface{} {
	return nil
}

// EmptyMap returns an empty (zero value) map
func EmptyMap() map[int]int {
	return nil
}

// EmptySlice returns an empty (zero value) slice
func EmptySlice() []int {
	return nil
}

// EmptyString returns an empty (zero value) string
func EmptyString() string {
	return ""
}

// EmptyChannel returns an empty (zero value) channel
func EmptyChannel() chan int {
	return nil
}

// EmptyPointer returns an empty (zero value) pointer
func EmptyPointer() *int {
	return nil
}

// EmptyBool returns an empty (zero value) bool
func EmptyBool() bool {
	return false
}

// EmptyFunc returns an empty (zero value) func
func EmptyFunc() func() {
	return nil
}

// EmptyInt returns an empty (zero value) int
func EmptyInt() int {
	return 0
}
