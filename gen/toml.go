package gen

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// getExcludedTestCases reads the file at tomlFilePath and manually parses the toml
// as no third party libraries are currently available (https://github.com/exercism/go/issues/2055).
func getExcludedTestCases(tomlFilePath string) (map[string]struct{}, error) {
	if _, err := os.Stat(tomlFilePath); err != nil {
		return nil, fmt.Errorf("tests.toml can't be found (%q)", err)
	}

	file, err := os.Open(tomlFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open %s: %w", tomlFilePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	excludedList := map[string]struct{}{}
	var currentCaseUUID string

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		switch {
		case line == "", strings.HasPrefix(line, "#"):
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
