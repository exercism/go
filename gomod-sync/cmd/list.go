package cmd

import (
	"fmt"

	"github.com/exercism/go/gomod-sync/gomod"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List go.mod files and the Go version they specify",
	RunE: func(cmd *cobra.Command, args []string) error {
		files, err := gomod.Infos(exercisesPathFlag)
		if err != nil {
			return fmt.Errorf("could not get go.mod information: %w", err)
		}

		for _, file := range files {
			fmt.Printf("%s => %s\n", file.Path, file.GoVersion)
		}

		return nil
	},
}
