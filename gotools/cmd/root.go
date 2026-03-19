// Package cmd contains the logic for the various sub-commands.
// root.go represents the main command and contains common flags and setup.
package cmd

import (
	"fmt"
	"os"

	"github.com/exercism/go/gotools/config"
	"github.com/spf13/cobra"
)

var (
	exercisesPathFlag string
	targetVersionFlag string
	configFileFlag    string
	// Should we make automated updates to resolve issues.
	updateFlag bool
	configData config.Config
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&exercisesPathFlag,
		"exercises", "e", "../exercises",
		"path to the exercises folder. go.mod files will "+
			"be recursively searched inside this directory.")
	rootCmd.PersistentFlags().StringVarP(&configFileFlag,
		"config", "c", "config.json",
		"path to the JSON configuration file")
}

var rootCmd = &cobra.Command{
	Use:   "gotools",
	Short: "gotools checks and updates the go version for all go.mod files in a path.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() {
	rootCmd.SilenceUsage = true
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func loadConfig(cmd *cobra.Command, args []string) error {
	var err error

	// Load version config
	configData, err = config.Load(configFileFlag)

	versionWasGiven := cmd.PersistentFlags().Changed("goversion")

	if err != nil && !versionWasGiven {
		return fmt.Errorf("failed to load config file and flag --goversion not present: %v", err)
	}

	// Override config default if version passed via flag
	if versionWasGiven {
		configData.GoModVersion.Default = targetVersionFlag
	}

	return nil
}
