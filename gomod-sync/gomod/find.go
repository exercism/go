package gomod

import (
	"fmt"
	"os"
	"path/filepath"
)

// FindFiles finds all the paths to go.mod files inside the given directory.
func FindFiles(dirPath string) ([]string, error) {
	var files []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if info.Name() == "go.mod" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("walking dir %q: %w", dirPath, err)
	}

	return files, err
}
