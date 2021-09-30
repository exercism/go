package logs

import "strings"

func Application(log string) string {
	runeToApp := make(map[rune]string)

	runeToApp['❗'] = "recommendation"
	runeToApp['🔍'] = "search"
	runeToApp['☀'] = "weather"

	for _, char := range log {
		if _, ok := runeToApp[char]; ok {
			return runeToApp[char]
		}
	}

	return "default"
}

func Redact(log string, redactions []rune) string {
	var result strings.Builder

	for _, char := range log {
		mustRedact := false

		for _, redaction := range redactions {
			if char == redaction {
				mustRedact = true
				break
			}
		}

		if mustRedact {
			continue
		}

		result.WriteRune(char)
	}

	return result.String()
}

func Replace(log string, old, new rune) string {
	var result strings.Builder

	for _, char := range log {
		if char == old {
			char = new
		}

		result.WriteRune(char)
	}

	return result.String()
}

func Count(log string) int {
	var count int

	for range log {
		count++
	}

	return count
}
