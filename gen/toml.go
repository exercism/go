package gen

import (
	"bufio"
	"os"
	"strings"
)

func getExcludedTestCases(tomlFilePath string) (map[string]struct{}, error) {
	file, err := os.Open(tomlFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	excludedList := map[string]struct{}{}
	var currentCaseUUID string

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		switch {
		case len(line) == 0, strings.HasPrefix(line, "#"):
			continue
		case strings.HasPrefix(line, "["):
			currentCaseUUID = strings.Trim(line, "[]")
		case isExcluded(line):
			excludedList[currentCaseUUID] = struct{}{}
		}
	}

	return excludedList, nil
}

func isExcluded(line string) bool {
	parts := strings.Split(line, "=")
	key := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])
	if key == "include" {
		return value == "false"
	}
	return false
}
