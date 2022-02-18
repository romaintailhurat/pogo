/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"net/http"
)

// hc command
var hcCmd = &cobra.Command{
	Use:   "hc",
	Short: "Health check on the current environment",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("https://pogues-api.kube.developpement.insee.fr/api/healthcheck")

		if err != nil {
			panic(err)
		}

		defer resp.Body.Close()

		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			fmt.Println("Service OK")
		} else {
			fmt.Println("Service KO")
			fmt.Println("Response status:", resp.Status)
		}

	},
}

func init() {
	rootCmd.AddCommand(hcCmd)
}
