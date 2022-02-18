/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Base command
var rootCmd = &cobra.Command{
	Use:   "pogo",
	Short: "Pogues CLI",
	Long: `The command line for interacting with the Pogues API.

A default environment is configured with the CLI, you can add some others.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
