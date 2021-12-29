package cmd

import (
	"fmt"

	"github.com/exercism/go/gomod-sync/gomod"
	"github.com/logrusorgru/aurora/v3"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	SilenceErrors:     true,
	Use:               "check",
	Short:             "Checks if all go.mod files are in the target version",
	PersistentPreRunE: loadConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		files, err := gomod.Infos(exercisesPathFlag)
		if err != nil {
			return err
		}

		type faultyFile struct {
			gomod.Info
			ExpectedVersion string
		}

		var faultyFiles []faultyFile
		for _, file := range files {
			expectedVersion := versionConfig.ExerciseExpectedVersion(file.ExerciseSlug)
			if file.GoVersion != expectedVersion {
				fmt.Println(aurora.Red(fmt.Sprintf("%v has version %s, but %s expected - FAIL", file.Path, file.GoVersion, expectedVersion)))
				faultyFiles = append(faultyFiles, faultyFile{Info: file, ExpectedVersion: expectedVersion})
			} else {
				fmt.Println(aurora.Green(fmt.Sprintf("%v has version %s as expected - OK", file.Path, file.GoVersion)))
			}
		}

		if len(faultyFiles) > 0 {
			fmt.Println(aurora.Red(fmt.Sprintf("The following %d go.mod file(s) do not have the correct version set:", len(faultyFiles))))
			for _, file := range faultyFiles {
				fmt.Println(aurora.Red(fmt.Sprintf("\t%v has version %s, but %s expected", file.Path, file.GoVersion, file.ExpectedVersion)))
			}
			return fmt.Errorf("%d go.mod file(s) are not in the target version", len(faultyFiles))
		}

		return nil
	},
}
