package cmd

import (
	"encoding/json"
	"os"

	"github.com/spf13/cobra"
)

var (
	exercisesPathFlag string
	configFileFlag string
	config Config
)

func init() {
	rootCmd.PersistentFlags().StringVarP(&exercisesPathFlag,
		"exercises", "e", "../../exercises",
		"path to the exercises folder. config.json files will "+
			"be recursively searched inside this directory.")
	rootCmd.PersistentFlags().StringVarP(&configFileFlag,
		"config", "c", "config.json",
		"path to the JSON configuration file. ")
}

var rootCmd = &cobra.Command{
	Use:   "editor-check",
	Short: "editor-check checks that all the test files are in the .meta/config.json editor",
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

type Config struct {
	Exclude map[string][]string `json:"exclude"`
}

func loadConfig(cmd *cobra.Command, args []string) error {
	var err error

	var c Config
	jsonData, err := os.ReadFile(configFileFlag)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, &c)
	if err != nil {
		return err
	}
	config = c
	return nil
}
