package cmd

import (
	"fmt"
	"os"

	"github.com/exercism/go/gomod-sync/cmd/config"
	"github.com/spf13/cobra"
)

var (
	exercisesPathFlag string
	targetVersionFlag string
	configFileFlag    string
	versionConfig     config.VersionConfig
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&exercisesPathFlag,
		"exercises", "e", "../exercises",
		"path to the exercises folder. go.mod files will "+
			"be recursively searched inside this directory.")
	rootCmd.PersistentFlags().StringVarP(&targetVersionFlag,
		"goversion", "v", "",
		"target go version that all go.mod files are expected to have. "+
			"This will be used to check if the go.mod files are in the expected "+
			"version in case of the check command, and to update all go.mod files to this version "+
			"in the case of the update command. Using this flag will override "+
			"the version specified in the config file.")
	rootCmd.PersistentFlags().StringVarP(&configFileFlag,
		"config", "c", "config.json",
		"path to the JSON configuration file. ")
}

var rootCmd = &cobra.Command{
	Use:   "gomod-sync",
	Short: "gomod-sync checks and updates the go version for all go.mod files in a path.",
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
	versionConfig, err = config.Load(configFileFlag)

	versionWasGiven := cmd.PersistentFlags().Changed("goversion")

	if err != nil && !versionWasGiven {
		return fmt.Errorf("failed to load config file and flag --goversion not present: %v", err)
	}

	// Override config default if version passed via flag
	if versionWasGiven {
		versionConfig.Default = targetVersionFlag
	}

	return nil
}
