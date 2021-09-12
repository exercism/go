package logs

// Message extracts the message from the provided log line.
func Message(line string) string {
	panic("Please implement the Message() function")
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	panic("Please implement the MessageLen() function")
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	panic("Please implement the LogLevel() function")
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	panic("Please implement the Reformat() function")
}
