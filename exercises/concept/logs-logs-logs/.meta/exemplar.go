package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	runeToApp := map[rune]string{
		'❗': "recommendation",
		'🔍': "search",
		'☀': "weather",
	}

	for _, char := range log {
		if _, ok := runeToApp[char]; ok {
			return runeToApp[char]
		}
	}

	return "default"
}

// Replace replaces all occurances of old with new, returning the modified log
// to the caller.
func Replace(log string, old, new rune) string {
	var modifiedLog string

	for _, char := range log {
		if char == old {
			char = new
		}

		modifiedLog += string(char)
	}

	return modifiedLog
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	runeCount := utf8.RuneCountInString(log)
	return runeCount <= limit
}
