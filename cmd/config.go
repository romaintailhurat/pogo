/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/romaintailhurat/pogo/types"
)

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
		}
		defer file.Close()

		config := types.GetConf()

		for _, env := range config.Envs {
			msg := fmt.Sprintf("%s → %s", env.Id, env.Url)
			fmt.Println(msg)
		}

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
