package bob

import "strings"

const testVersion = 3

// Hey returns Bob's responses to a given remark.
func Hey(remark string) string {
	switch remark = strings.TrimSpace(remark); {
	case silent(remark):
		return "Fine. Be that way!"
	case yelling(remark):
		return "Whoa, chill out!"
	case asking(remark):
		return "Sure."
	default:
		return "Whatever."
	}
}

func yelling(remark string) bool {
	return strings.ToUpper(remark) == remark && strings.ToLower(remark) != strings.ToUpper(remark)
}

func asking(remark string) bool {
	return strings.HasSuffix(remark, "?")
}

func silent(remark string) bool {
	return remark == ""
}
