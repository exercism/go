package config

import "sort"

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

// Equal tells if this VersionConfig is equal to another VersionConfig
func (vc *VersionConfig) Equal(other VersionConfig) bool {
	return vc.Default == other.Default && equalExceptions(vc.Exceptions, other.Exceptions)
}

// equalExceptions compares two lists of exercise versions and tells if they are equal.
// Two exercise list versions are considered equal if they contain the same elements,
// regardless of their order in both lists.
func equalExceptions(a, b []ExerciseVersion) bool {
	if len(a) != len(b) {
		return false
	}

	// Copy slices
	aCopy := make([]ExerciseVersion, len(a))
	bCopy := make([]ExerciseVersion, len(b))
	copy(aCopy, a)
	copy(bCopy, b)
	sort.Slice(aCopy, func(i, j int) bool { return aCopy[i].Exercise < aCopy[j].Exercise })
	sort.Slice(bCopy, func(i, j int) bool { return bCopy[i].Exercise < bCopy[j].Exercise })

	for i := range aCopy {
		if aCopy[i] != bCopy[i] {
			return false
		}
	}

	return true
}

// ExerciseVersion represents the version one particular exercise should
// have in its go.mod file
type ExerciseVersion struct {
	Exercise string `json:"exercise"`
	Version  string `json:"version"`
}
