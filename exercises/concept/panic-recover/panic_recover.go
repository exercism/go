package panicrecover

// Add condition to panic
func AccessNamesPanic(name []string, index int) string {
	panic("Implement AddPanic() method")
}

func RecoverPanic() (msg interface{}) {
	panic("Implement RecoverPanic() method")
}

// Resolve error causing panic.
// Return the last element of the names slice without causing panic
func ResolveError(names []string, index int) string {
	panic("Implement ResolveError() Method")
}
