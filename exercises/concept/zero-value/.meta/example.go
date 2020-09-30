package _meta

// EmptyBool returns an empty (zero value) bool
func EmptyBool() bool {
	return false
}

// EmptyInt returns an empty (zero value) int
func EmptyInt() int {
	return 0
}

// EmptyString returns an empty (zero value) string
func EmptyString() string {
	return ""
}

// EmptyFunc returns an empty (zero value) func
func EmptyFunc() func() {
	return nil
}

// EmptyPointer returns an empty (zero value) pointer
func EmptyPointer() *int {
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

// EmptyChannel returns an empty (zero value) channel
func EmptyChannel() chan int {
	return nil
}

// EmptyInterface returns an empty (zero value) interface
func EmptyInterface() interface{} {
	return nil
}
