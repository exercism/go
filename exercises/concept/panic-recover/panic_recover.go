package panicrecover

// Add condition to panic
func AddPanic(name []string, index int) string {
	defer RecoverPanic()
	val := name[index]
	return val
}

func RecoverPanic() (msg interface{}) {
	msg = recover()
	if msg != nil {
		return msg
	}
	return
}

// Resolve error causing panic. Return the last element of the names slice without causing panic
func ResolveError(names []string, index int) string {
	len := len(names)
	val := names[len-1]
	return val
}
