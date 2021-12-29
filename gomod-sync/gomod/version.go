package gomod

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

var versionRegex *regexp.Regexp

func init() {
	// versionRegex is a regex that matches the Go version
	// information in a go.mod file
	versionRegex = regexp.MustCompile(`go (\d+\.\d+)`)
}

// GoVersion fetches the Go version specified by a go.mod file.
func GoVersion(gomodFilePath string) (string, error) {
	// Open gomodFilePath for reading
	file, err := os.Open(gomodFilePath)
	if err != nil {
		return "", fmt.Errorf("error opening %q: %w", gomodFilePath, err)
	}
	defer file.Close()

	// Read go.mod file
	rawContent, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error reading %q: %w", gomodFilePath, err)
	}

	// Fetch version
	content := string(rawContent)
	version := versionRegex.FindStringSubmatch(content)
	if len(version) < 2 {
		return "", fmt.Errorf("could not find version info in %q", gomodFilePath)
	}

	return version[1], nil
}

// Update updates a go.mod file at the given path to the given Go version.
func Update(path string, version string) error {
	// Open path file for writing
	file, err := os.OpenFile(path, os.O_RDWR, 0o644)
	if err != nil {
		return fmt.Errorf("error opening %q: %w", path, err)
	}
	defer file.Close()

	// Read the file
	rawContent, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading %q: %w", path, err)
	}

	// Replace the version
	content := string(rawContent)
	content = versionRegex.ReplaceAllString(content, fmt.Sprintf("go %s", version))

	// Truncate the file
	if err = file.Truncate(0); err != nil {
		return fmt.Errorf("error truncating %q: %w", path, err)
	}
	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("error seeking %q: %w", path, err)
	}

	// Write the file
	_, err = file.WriteString(content)
	if err != nil {
		return fmt.Errorf("error writing %q: %w", path, err)
	}
	return nil
}
