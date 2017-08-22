package bob

import "strings"

// Hey returns Bob's responses to any given dialogue
func Hey(dialogue string) string {
	switch dialogue = strings.TrimSpace(dialogue); {
	case silent(dialogue):
		return "Fine. Be that way!"
	case yelling(dialogue):
		return "Whoa, chill out!"
	case asking(dialogue):
		return "Sure."
	default:
		return "Whatever."
	}
}

func yelling(dialogue string) bool {
	return strings.ToUpper(dialogue) == dialogue && strings.ToLower(dialogue) != strings.ToUpper(dialogue)
}

func asking(dialogue string) bool {
	return strings.HasSuffix(dialogue, "?")
}

func silent(dialogue string) bool {
	return dialogue == ""
}

// This test versioning is specific to Exercism,
// you don't need to worry about this now.
const testVersion = 3
