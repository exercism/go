package gomod_test

import (
	"path/filepath"
	"testing"

	"github.com/exercism/go/gomod-sync/gomod"
)

func TestInfos(t *testing.T) {
	tests := []struct {
		Name        string
		BasePath    string
		Expected    []gomod.Info
		ExpectError bool
	}{
		{
			Name:     "Find go mod information in folder with 2 exercises",
			BasePath: filepath.Join("..", "testdata"),
			Expected: []gomod.Info{
				{
					Path:         filepath.Join("..", "testdata", "exercise01", "go.mod"),
					GoVersion:    "1.17",
					ExerciseSlug: "exercise01",
				},
				{
					Path:         filepath.Join("..", "testdata", "exercise02", "go.mod"),
					GoVersion:    "1.17",
					ExerciseSlug: "exercise02",
				},
			},
			ExpectError: false,
		},
		{
			Name:        "Find go mod information in non-existent folder",
			BasePath:    filepath.Join("..", "testdata", "non-existent"),
			Expected:    nil,
			ExpectError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			infos, err := gomod.Infos(test.BasePath)
			if test.ExpectError && err == nil {
				t.Fatalf("expected error, but got none")
			}

			if !test.ExpectError && err != nil {
				t.Fatalf("didn't expect error, but got %v", err)
			}

			if !infoSlicesEqual(infos, test.Expected) {
				t.Fatalf("expected infos: %v, got %v", test.Expected, infos)
			}
		})
	}
}

func infoSlicesEqual(a, b []gomod.Info) bool {
	if len(a) != len(b) {
		return false
	}

	if len(a) == 0 {
		return true
	}

	size := len(a)
	for i := 0; i < size; i++ {
		if !equalInfos(a[i], b[i]) {
			return false
		}
	}

	return true
}

func equalInfos(a, b gomod.Info) bool {
	return a.Path == b.Path &&
		a.GoVersion == b.GoVersion &&
		a.ExerciseSlug == b.ExerciseSlug
}
