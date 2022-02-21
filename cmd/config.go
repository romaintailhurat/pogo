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

type Env struct {
	id  string
	url string
}

type Config struct {
	envs []Env
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
		}
		defer file.Close()

		baseEnv := Env{id: "kube", url: "https://pogues.kube.developpement.insee.fr/"}
		defaultConfig := Config{envs: []Env{baseEnv}}

		for _, env := range defaultConfig.envs {
			msg := fmt.Sprintf("%s → %s", env.id, env.url)
			fmt.Println(msg)
		}

	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
