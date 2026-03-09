package microblog

func Truncate(phrase string) string {
	runes := []rune(phrase)
	return string(runes[:min(len(runes), 5)])
}
