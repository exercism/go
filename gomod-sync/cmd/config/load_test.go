package config_test

import (
	"path/filepath"
	"testing"

	"github.com/exercism/go/gomod-sync/cmd/config"
)

func TestLoad(t *testing.T) {
	tests := []struct {
		Name        string
		Path        string
		Expected    config.VersionConfig
		ExpectError bool
	}{
		{
			Name:        "Loading non-existent config",
			Path:        filepath.Join("..", "..", "testdata", "non_existent.json"),
			Expected:    config.VersionConfig{},
			ExpectError: true,
		},
		{
			Name: "Loading config with no exceptions",
			Path: filepath.Join("..", "..", "testdata", "version_config_no_exceptions.json"),
			Expected: config.VersionConfig{
				Default: "1.16",
			},
			ExpectError: false,
		},
		{
			Name: "Loading config with 1 exception",
			Path: filepath.Join("..", "..", "testdata", "version_config_one_exception.json"),
			Expected: config.VersionConfig{
				Default: "1.16",
				Exceptions: []config.ExerciseVersion{
					{
						Exercise: "exercise01",
						Version:  "1.17",
					},
				},
			},
			ExpectError: false,
		},
		{
			Name: "Loading config with 2 exceptions",
			Path: filepath.Join("..", "..", "testdata", "version_config_two_exceptions.json"),
			Expected: config.VersionConfig{
				Default: "1.16",
				Exceptions: []config.ExerciseVersion{
					{
						Exercise: "exercise01",
						Version:  "1.17",
					},
					{
						Exercise: "exercise02",
						Version:  "1.17",
					},
				},
			},
			ExpectError: false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			config, err := config.Load(test.Path)

			if test.ExpectError && err == nil {
				t.Fatalf("expected error, but got none")
			}

			if !test.ExpectError && err != nil {
				t.Fatalf("didn't expect error, but got %v", err)
			}

			if !configEqual(config, test.Expected) {
				t.Fatalf("expected config %+v, but got %+v", test.Expected, config)
			}
		})
	}
}

func configEqual(a, b config.VersionConfig) bool {
	return a.Default == b.Default && equalExceptions(a.Exceptions, b.Exceptions)
}

// equalExceptions compares two lists of exercise versions and tells if they are equal.
// Two exercise list versions are considered equal if they contain the same elements
// in the same order
func equalExceptions(a, b []config.ExerciseVersion) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
