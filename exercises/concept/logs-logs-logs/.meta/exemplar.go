package logs

import (
	"fmt"
	"strings"
)

// Message extracts the message from the provided log line.
func Message(line string) string {
	parts := strings.Split(line, ":")
	if len(parts) < 2 {
		return ""
	}

	return strings.TrimSpace(parts[1])
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	return len([]rune(Message(line)))
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	part1 := strings.Split(line, ":")[0]
	logLevel := strings.Trim(part1, "[]")
	return strings.ToLower(logLevel)
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	return fmt.Sprintf("%s (%s)", Message(line), LogLevel(line))
}
