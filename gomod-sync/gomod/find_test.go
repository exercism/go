package gomod_test

import (
	"path/filepath"
	"testing"

	"github.com/exercism/go/gomod-sync/gomod"
)

func TestFindFiles(t *testing.T) {
	tests := []struct {
		Name        string
		BasePath    string
		Expected    []string
		ExpectError bool
	}{
		{
			Name:     "Find go.mod files in directory with 2 exercises",
			BasePath: filepath.Join("..", "testdata"),
			Expected: []string{
				filepath.Join("..", "testdata", "exercise01", "go.mod"),
				filepath.Join("..", "testdata", "exercise02", "go.mod"),
			},
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			files, err := gomod.FindFiles(test.BasePath)

			if test.ExpectError && err == nil {
				t.Fatalf("expected error, but got none")
			}

			if !test.ExpectError && err != nil {
				t.Fatalf("didn't expect error, but got %v", err)
			}

			if !stringSlicesEqual(test.Expected, files) {
				t.Fatalf("expected %v go.mod files, found %v", test.Expected, files)
			}
		})
	}
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if len(a) == 0 {
		return true
	}

	size := len(a)
	for i := 0; i < size; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
