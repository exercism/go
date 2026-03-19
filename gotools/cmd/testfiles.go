// testfiles.go contains the logic for the test_files subcommand, used to check the test.files values in exercise .meta/config.json files.
package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	aurora "github.com/logrusorgru/aurora/v3"
	"github.com/spf13/cobra"
)

// Patterns to strip from JSON output.
var EmptyJson = []string{
	`"[^"]*":null,`,
	`"[^"]*":\[\],`,
	`"[^"]*":"",`,
	`,"[^"]*":null`,
	`,"[^"]*":\[\]`,
	`,"[^"]*":""`,
}

func init() {
	rootCmd.AddCommand(testFilesCmd)
	testFilesCmd.PersistentFlags().BoolVarP(&updateFlag,
		"update", "u", false,
		"make automated updates to resolve issues")
}

type ExerciseConfigFile struct {
	Authors      []string `json:"authors"`
	Contributors []string `json:"contributors"`
	Files        struct {
		Solution    []string `json:"solution"`
		Test        []string `json:"test"`
		Example     []string `json:"example"`
		Editor      []string `json:"editor"`
		Invalidator []string `json:"invalidator"`
	} `json:"files"`
	Blurb      string `json:"blurb"`
	Source     string `json:"source"`
	Source_url string `json:"source_url"`
}

type ExerciseData struct {
	Slug string
	Type string
	Path string
}

func (e ExerciseData) Check() error {
	configFile := filepath.Join(e.Path, ".meta", "config.json")
	jsonData, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}
	var meta ExerciseConfigFile
	err = json.Unmarshal(jsonData, &meta)
	if err != nil {
		return err
	}
	slices.Sort(meta.Files.Test)

	solutionFile := strings.ReplaceAll(e.Slug, "-", "_") + ".go"
	srcFiles, err := filepath.Glob(filepath.Join(e.Path, "*.go"))
	if err != nil {
		return err
	}
	var testFiles, editorFiles []string
	for _, srcFile := range srcFiles {
		_, name := filepath.Split(srcFile)
		if name != solutionFile && !slices.Contains(configData.TestFiles.Exclude[e.Slug], name) {
			if strings.Contains(name, "test") {
				testFiles = append(testFiles, name)
			} else {
				editorFiles = append(editorFiles, name)
			}
		}
	}
	slices.Sort(testFiles)
	slices.Sort(editorFiles)

	if !slices.Equal(meta.Files.Test, testFiles) || !slices.Equal(meta.Files.Editor, editorFiles) {
		if updateFlag {
			meta.Files.Editor = editorFiles
			meta.Files.Test = testFiles
			// Convert to JSON, remove empty fields and indent.
			out, err := json.Marshal(meta)
			for _, pattern := range EmptyJson {
				out = regexp.MustCompile(pattern).ReplaceAll(out, []byte{})
			}

			var indented bytes.Buffer
			json.Indent(&indented, out, "", "  ")
			indented.WriteString("\n")

			if err != nil {
				return fmt.Errorf("exercise %q, failed to marshal config data, %v", e.Slug, err)
			}
			fi, err := os.Stat(configFile)
			if err != nil {
				return fmt.Errorf("exercise %q, failed to stat() config data, %v", e.Slug, err)
			}
			if err = os.WriteFile(configFile, indented.Bytes(), fi.Mode().Perm()); err != nil {
				return fmt.Errorf("exercise %q, failed to write config data, %v", e.Slug, err)
			}
			fmt.Println(aurora.Green(fmt.Sprintf("Updated the config file for %q", e.Slug)))
		} else {
			return fmt.Errorf("exercise %q, config files do not match source files. Meta files test [%v] and editor [%v] vs source test [%v] and others [%v]", e.Slug, meta.Files.Test, meta.Files.Editor, testFiles, editorFiles)
		}
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

var testFilesCmd = &cobra.Command{
	SilenceErrors:     true,
	Use:               "test_files",
	Short:             "Checks all test files are shown in the test tab",
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
