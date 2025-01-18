package config_test

import (
	"testing"

	"github.com/exercism/go/gomod-sync/cmd/config"
)

func TestExpectedVersion(t *testing.T) {
	tests := []struct {
		Name     string
		Config   config.VersionConfig
		Exercise string
		Expected string
	}{
		{
			Name: "Config with just a default version should return the default version",
			Config: config.VersionConfig{
				Default:    "1.16",
				Exceptions: nil,
			},
			Exercise: "exercise01",
			Expected: "1.16",
		},
		{
			Name: "Expected version for an exercise not in exceptions should return the default version",
			Config: config.VersionConfig{
				Default: "1.16",
				Exceptions: []config.ExerciseVersion{
					{
						Exercise: "exercise02",
						Version:  "1.17",
					},
				},
			},
			Exercise: "exercise01",
			Expected: "1.16",
		},
		{
			Name: "Expected version for an exercise in exceptions should return the version in the exception",
			Config: config.VersionConfig{
				Default: "1.16",
				Exceptions: []config.ExerciseVersion{
					{
						Exercise: "exercise02",
						Version:  "1.17",
					},
				},
			},
			Exercise: "exercise02",
			Expected: "1.17",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			got := test.Config.ExerciseExpectedVersion(test.Exercise)
			if got != test.Expected {
				t.Fatalf("expected version %q, got %q", test.Expected, got)
			}
		})
	}
}
