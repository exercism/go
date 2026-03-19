// gomod.go contains the `gomod` subcommand logic, used to check the go version in go.mod files.
package cmd

import (
	"fmt"

	aurora "github.com/logrusorgru/aurora/v3"
	"github.com/spf13/cobra"

	"github.com/exercism/go/gotools/gomod"
)

func init() {
	rootCmd.AddCommand(gomodCmd)
	gomodCmd.PersistentFlags().BoolVarP(&updateFlag,
		"update", "u", false,
		"make automated updates to resolve issues")
	gomodCmd.PersistentFlags().StringVarP(&targetVersionFlag,
		"goversion", "v", "",
		"target go version that all go.mod files are expected to have. "+
			"This will be used to check if the go.mod files are in the expected "+
			"version in case of the gomod command, and to update all go.mod files to this version "+
			"in the case of the update command. Using this flag will override "+
			"the version specified in the config file.")
}

var gomodCmd = &cobra.Command{
	SilenceErrors:     true,
	Use:               "gomod",
	Short:             "Checks if all go.mod files are in the target version",
	PersistentPreRunE: loadConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		files, err := gomod.Infos(exercisesPathFlag)
		if err != nil {
			return err
		}

		type faultyFile struct {
			Info            gomod.Info
			ExpectedVersion string
		}

		var faultyFiles []faultyFile
		for _, file := range files {
			expectedVersion := configData.GoModVersion.ExerciseExpectedVersion(file.ExerciseSlug)
			if file.GoVersion == expectedVersion {
				fmt.Println(aurora.Green(fmt.Sprintf("%v has version %s as expected - OK", file.Path, file.GoVersion)))
			} else {
				if updateFlag {
					if err := gomod.Update(file.Path, expectedVersion); err != nil {
						fmt.Println(aurora.Red(fmt.Sprintf("failed to update %q: %v", file.Path, err)))
						return fmt.Errorf("failed to update %q: %w", file.Path, err)
					} else {
						fmt.Println(aurora.Green(fmt.Sprintf("Updated %s: %s => %s\n", file.Path, file.GoVersion, expectedVersion)))
					}
				} else {
					fmt.Println(aurora.Red(fmt.Sprintf("%v has version %s, but %s expected - FAIL", file.Path, file.GoVersion, expectedVersion)))
					faultyFiles = append(faultyFiles, faultyFile{Info: file, ExpectedVersion: expectedVersion})
				}
			}
		}

		if len(faultyFiles) > 0 {
			fmt.Println(aurora.Red(fmt.Sprintf("The following %d go.mod file(s) do not have the correct version set:", len(faultyFiles))))
			for _, file := range faultyFiles {
				fmt.Println(aurora.Red(fmt.Sprintf("\t%v has version %s, but %s expected", file.Info.Path, file.Info.GoVersion, file.ExpectedVersion)))
			}
			return fmt.Errorf("%d go.mod file(s) are not in the target version", len(faultyFiles))
		}

		return nil
	},
}
