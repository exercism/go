package panicrecover

// Add condition to panic
func AddPanic(name []string, index int) string {
	defer RecoverPanic()
	val := name[index]
	return val
}

func RecoverPanic() (message interface{}) {
	message = recover()
	if message != nil {
		return message
	}
	return
}

// Resolve error causing panic
func ResolveError(names []string, index int) string {
	len := len(names)
	val := names[len-1]
	return val
}
