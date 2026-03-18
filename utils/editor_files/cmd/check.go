package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	aurora "github.com/logrusorgru/aurora/v3"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

type ExerciseConfigFile struct {
	Files struct {
		Editor []string `json:"editor"`
	} `json:"files"`
}

type ExerciseData struct {
	Slug string
	Type string
	Path string
}

func (e ExerciseData) Check() error {
	jsonData, err := os.ReadFile(filepath.Join(e.Path, ".meta", "config.json"))
	if err != nil {
		return err
	}
	var meta ExerciseConfigFile
	err = json.Unmarshal(jsonData, &meta)
	if err != nil {
		return err
	}
	slices.Sort(meta.Files.Editor)

	solutionFile := strings.ReplaceAll(e.Slug, "-", "_") + ".go"
	srcFiles, err := filepath.Glob(filepath.Join(e.Path, "*.go"))
	if err != nil {
		return err
	}
	var srcNames []string
	for _, srcFile := range srcFiles {
		_, name := filepath.Split(srcFile)
		if name != solutionFile && !slices.Contains(config.Exclude[e.Slug], name) {
			srcNames = append(srcNames, name)
		}
	}
	slices.Sort(srcNames)

	if !slices.Equal(meta.Files.Editor, srcNames) {
		return fmt.Errorf("exercise %q, editor files %v do not match source files %v", e.Slug, meta.Files.Editor, srcNames)
	}
	return nil
}

func exerciseData(exercisePath string) ([]ExerciseData, error) {
	var exercises []ExerciseData

	matches, err := filepath.Glob(filepath.Join(exercisePath, "practice", "*"))
	if err != nil {
		return nil, err
	}
	for _, match := range matches {
		file, err := os.Stat(match)
		if err != nil {
			return nil, err
		}
		if !file.IsDir() || file.Name()[0] == '.' {
			continue
		}

		parent, slug := filepath.Split(match)
		_, exerciseType := filepath.Split(parent)
		exercises = append(exercises, ExerciseData{Slug: slug, Type: exerciseType, Path: match})
	}
	return exercises, nil
}

var checkCmd = &cobra.Command{
	SilenceErrors:     true,
	Use:               "check",
	Short:             "Checks all test files are shown in the editor",
	PersistentPreRunE: loadConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		exercises, err := exerciseData(exercisesPathFlag)
		if err != nil {
			return err
		}
		var errs error
		for _, exercise := range exercises {
			if err := exercise.Check(); err != nil {
				fmt.Println(aurora.Red(err))
				errs = errors.Join(errs, err)
			}
		}
		return errs
	},
}
