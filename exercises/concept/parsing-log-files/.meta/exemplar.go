package parsing_log_files

import (
	"regexp"
)

func IsValidLine(text string) bool {
	str := `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`
	result, ok := regexp.MatchString(str, text)
	if ok != nil {
		panic("invalid regex: " + str)
	}
	return result
}

func SplitLogLine(text string) []string {
	str := `<[*~=-]*>`
	re := regexp.MustCompile(str)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
	str := `(?i)".*password.*"`
	regex := regexp.MustCompile(str)
	for _, line := range lines {
		if regex.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	regex := regexp.MustCompile(`end-of-line\d+`)
	return regex.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var result []string
	regex := regexp.MustCompile(`\sUser\s+(\S+)`)
	for _, line := range lines {
		match := regex.FindStringSubmatch(line)
		if len(match) > 1 {
			line = "[USR] " + match[1] + " " + line
		}
		result = append(result, line)
	}
	return result
}
