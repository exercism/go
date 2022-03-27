package go/exercises/concept/panic-recover

// Add condition to panic
func AddPanic(names []string, index int) {
	panic("Add panic with message :- `Index Out Of Bounds`")
}

// Resolve error causing panic
func ResoveError() {
	panic("Please resolve the error first")
}

// Recover from panic
func RecoverPanic() {
	panic("Please add the recovering logic for the panic caused by Index Out of Bounds Error")
}
