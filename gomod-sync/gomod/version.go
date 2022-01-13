package gomod

import (
	"fmt"
	"os"
	"regexp"
)

var versionRegex *regexp.Regexp

func init() {
	// versionRegex is a regex that matches the Go version
	// information in a go.mod file
	versionRegex = regexp.MustCompile(`(?m:^(\s*)go (\d+\.\d+)(\s*)$)`)
}

// GoVersion fetches the Go version specified by a go.mod file.
func GoVersion(gomodFilePath string) (string, error) {
	// Read go.mod file
	rawContent, err := os.ReadFile(gomodFilePath)
	if err != nil {
		return "", fmt.Errorf("error reading %q: %w", gomodFilePath, err)
	}

	// Fetch version
	content := string(rawContent)
	version := versionRegex.FindStringSubmatch(content)
	if len(version) < 3 {
		return "", fmt.Errorf("could not find version info in %q", gomodFilePath)
	}

	return version[2], nil
}

// Update updates a go.mod file at the given path to the given Go version.
func Update(path string, version string) error {
	// Read the file
	rawContent, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading %q: %w", path, err)
	}

	// Replace the version
	content := string(rawContent)
	content = versionRegex.ReplaceAllString(content, "${1}go "+version+"${3}")

	// Write the file
	err = os.WriteFile(path, []byte(content), 0o644)
	if err != nil {
		return fmt.Errorf("error writing %q: %w", path, err)
	}
	return nil
}
