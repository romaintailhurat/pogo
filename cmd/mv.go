/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// mvCmd represents the mv command
var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "Moving questionnaires from one env to the other",
	Long: `Moving questionnaires from one env to the other:
	
	pogo mv questid source_env_id target_env_id
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// FIXME use flags for source and target envs ?
		fmt.Printf("Moving questionnaire %s from %s to %s \n", args[0], args[1], args[2])
	},
}

func init() {
	rootCmd.AddCommand(mvCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mvCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mvCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
