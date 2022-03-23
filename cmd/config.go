/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
		file, err := ioutil.ReadFile(userHomeDir + "/.pog/config.json")

		var config types.Config

		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("No config file, using defaults")
			config = types.GetConf()
		} else {
			fmt.Println("Using custom conf file")
			json.Unmarshal([]byte(file), &config)
		}

		for _, env := range config.Envs {
			msg := fmt.Sprintf("%s → %s", env.Id, env.Url)
			fmt.Println(msg)
		}

	},
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a config file",
	Long:  "Create a config file",
	Run: func(cmd *cobra.Command, args []string) {
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		fileLocation := userHomeDir + "/.pog/config.json"
		fmt.Println("Creating config file: " + fileLocation)

		os.MkdirAll(userHomeDir+"/.pog/", os.ModePerm)

		f, err := os.Create(fileLocation)

		if err != nil {
			panic(err)
		}

		j, err := json.Marshal(types.GetConf())

		if err != nil {
			panic(err)
		}

		lines, err := f.Write(j)

		if err != nil {
			f.Close()
			panic(err)
		}

		fmt.Sprintln(lines, " lines written to file")

		f.Close()

		//ioutil.WriteFile("test.json", j, 0755)
	},
}

func init() {
	configCmd.AddCommand(createCmd)
	rootCmd.AddCommand(configCmd)
}
