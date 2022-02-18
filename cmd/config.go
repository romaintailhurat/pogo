/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Config struct {
	envs string
}

// config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "pog configuration",
	Long: `pog comes with a default in memory configuration.

For custom configuration, either:

- create a $HOME/.pog/config.json
- use pog config create

The latter will create a config.json with the default configuration.

You will then be able to override it.
`,
	Run: func(cmd *cobra.Command, args []string) {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		file, err := os.Open(userHomeDir + "/.pog/config.json")
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("No config file, using defaults")
			fmt.Println("  To create a config file, use pog config create (WIP)")
		}
		defer file.Close()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
