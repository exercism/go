package panicrecover

// Add condition to panic
func AccessNames(name []string, index int) string {
	defer RecoverPanic()
	val := name[index]
	return val
}

func RecoverPanic() (errorMessage interface{}) {
	errorMessage = recover()
	if errorMessage != nil {
		return errorMessage
	}
	return
}

// Resolve error causing panic
func ResolveError(names []string, index int) string {
	lengthOfNames := len(names)
	val := names[lengthOfNames-1]
	return val
}
