package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	panic("Please implement the Application() function")
}

// Redact removes all occurances of the runes in redactions from the given log,
// returning the redacted log to the caller.
func Redact(log string, redactions []rune) string {
	panic("Please implement the Redact() function")
}

// Replace replaces all occurances old with new, returning the modified log to
// the caller.
func Replace(log string, old, new rune) string {
	panic("Please implement the Replace() function")
}

// Count determines the number of characters in log.
func Count(log string) int {
	panic("Please implement the Count() function")
}
