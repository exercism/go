package cmd

import (
	"fmt"

	"github.com/exercism/go/gomod-sync/gomod"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
}

var updateCmd = &cobra.Command{
	Use:               "update",
	Short:             "Updates go.mod files to the target version",
	PersistentPreRunE: loadConfig,
	RunE: func(cmd *cobra.Command, args []string) error {
		files, err := gomod.Infos(exercisesPathFlag)
		if err != nil {
			return fmt.Errorf("could not get go.mod information: %w", err)
		}

		for _, file := range files {
			expectedVersion := versionConfig.ExerciseExpectedVersion(file.ExerciseSlug)
			if file.GoVersion != expectedVersion {
				if err := gomod.Update(file.Path, expectedVersion); err != nil {
					return fmt.Errorf("failed to update %q: %w", file.Path, err)
				}
				fmt.Printf("Updated %s: %s => %s\n", file.Path, file.GoVersion, expectedVersion)
			}
		}

		return nil
	},
}
