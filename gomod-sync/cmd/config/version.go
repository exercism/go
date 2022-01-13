package config

// ExerciseVersion represents the version one particular exercise should
// have in its go.mod file
type ExerciseVersion struct {
	Exercise string `json:"exercise"`
	Version  string `json:"version"`
}

// VersionConfig represents the version configuration
type VersionConfig struct {
	// Default version for the go.mod files
	Default string `json:"default"`
	// List of exercises and their versions that must not have
	// the default version.
	Exceptions []ExerciseVersion `json:"exceptions"`
}

// ExerciseExpectedVersion gives the expected version the go.mod file for
// an exercise should be at. The argument should be the slug of an exercise.
func (vc *VersionConfig) ExerciseExpectedVersion(exercise string) string {
	for _, exception := range vc.Exceptions {
		if exception.Exercise == exercise {
			return exception.Version
		}
	}
	return vc.Default
}
