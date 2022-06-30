package gen

import (
	"bufio"
	"os"
	"strings"
)

func getExcludedTestCases(tomlFilePath string) (map[string]bool, error) {
	file, err := os.Open(tomlFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	excludedList := map[string]bool{}
	var currentCaseUUID string

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		length := len(line)
		if length == 0 {
			continue
		}
		switch string(line[0]) {
		case "#":
			continue
		case "[":
			currentCaseUUID = string(line[1 : length-1])
		default:
			isIncludeStmt, isCaseIncluded := verifyIncludeStatement(line)
			if isIncludeStmt && !isCaseIncluded {
				excludedList[currentCaseUUID] = true
			}
		}
	}

	return excludedList, nil
}

func verifyIncludeStatement(line string) (bool, bool) {
	parts := strings.Split(line, "=")
	key := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])
	if key == "include" {
		return true, value == "true"
	}
	return false, false
}
