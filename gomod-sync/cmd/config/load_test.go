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
			Name: "Loading non-existent config",
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
						Exercise: "exercise02",
						Version:  "1.17",
					},
					{
						Exercise: "exercise01",
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

			if !config.Equal(test.Expected) {
				t.Fatalf("expected config %+v, but got %+v", test.Expected, config)
			}
		})
	}
}
