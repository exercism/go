package gomod

import (
	"fmt"
	"path/filepath"
)

// Info represents the information of a go.mod file.
type Info struct {
	// Path to the go.mod file
	Path string
	// Go version specified by the go.mod file
	GoVersion string
	// Exercise slug
	ExerciseSlug string
}

// Infos returns information about all the go.mod files
// inside a given directory.
func Infos(dirPath string) (info []Info, err error) {
	var infos []Info

	files, err := FindFiles(dirPath)
	if err != nil {
		return nil, fmt.Errorf("finding go.mod infos: %w", err)
	}

	for _, file := range files {
		version, err := GoVersion(file)
		if err != nil {
			return nil, fmt.Errorf("could not get version for file %q: %w", file, err)
		}
		infos = append(infos, Info{
			Path:         file,
			GoVersion:    version,
			ExerciseSlug: filepath.Base(filepath.Dir(file)),
		})

	}

	return infos, nil
}
