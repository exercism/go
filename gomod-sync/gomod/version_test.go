package gomod_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/exercism/go/gomod-sync/gomod"
)

func TestGoVersion(t *testing.T) {
	tests := []struct {
		Name        string
		Path        string
		Expected    string
		ExpectError bool
	}{
		{
			Name:        "Get go version of go.mod with version 1.17",
			Path:        filepath.Join("..", "testdata", "exercise01", "go.mod"),
			Expected:    "1.17",
			ExpectError: false,
		},
		{
			Name:        "Get go version of non-existent go.mod",
			Path:        filepath.Join("..", "testdata", "non-existent", "go.mod"),
			Expected:    "",
			ExpectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			version, err := gomod.GoVersion(test.Path)
			if test.ExpectError && err == nil {
				t.Fatalf("expected error, but got none")
			}

			if !test.ExpectError && err != nil {
				t.Fatalf("didn't expect error, but got %v", err)
			}

			if version != test.Expected {
				t.Fatalf("expected version: %q, got %q", test.Expected, version)
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	tests := []struct {
		Name     string
		Version  string
		Original string
		Modified string
	}{
		{
			Name:    "Update version from 1.17 to 1.18",
			Version: "1.18",
			Original: `
module testmodule

go 1.17
`,
			Modified: `
module testmodule

go 1.18
`,
		},
	}

	temporaryDir := filepath.Join("..", "testdata", "tmp")

	err := os.MkdirAll(temporaryDir, 0o755)
	if err != nil {
		t.Fatalf("could not create temporary directory for test: %s", err.Error())
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			filePath := filepath.Join(temporaryDir, "go.mod")
			err := os.WriteFile(filePath, []byte(test.Original), 0o644)
			if err != nil {
				t.Fatalf("could not create temporary go.mod file for test: %s", err.Error())
			}

			err = gomod.Update(filePath, test.Version)
			if err != nil {
				t.Fatalf("could not update version of temporary go.mod for test: %s", err.Error())
			}

			data, err := os.ReadFile(filePath)
			if err != nil {
				t.Fatalf("could not read temporary go.mod for test: %s", err.Error())
			}

			if string(data) != test.Modified {
				t.Fatalf("expected go.mod file with the following contents:\n"+
					"%s\n"+
					"but got:\n"+
					"%s", test.Modified, string(data))
			}

			err = os.Remove(filePath)
			if err != nil {
				t.Fatalf("could not remove temporary go.mod file for test: %s", err.Error())
			}
		})
	}

	err = os.RemoveAll(temporaryDir)
	if err != nil {
		t.Fatalf("could not remove temporary directory for test: %s", err.Error())
	}
}
