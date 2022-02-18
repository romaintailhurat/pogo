/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

type Questionnaire struct {
	Name string
}

// get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the JSON representation of a questionnaire whose id is passed.",
	Long: `Get the JSON representation of a questionnaire whose id is passed.
	
	(WIP) The JSON itself is only returned to stdout.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		resp, err := http.Get("https://pogues-api.kube.developpement.insee.fr/api/persistence/questionnaire/" + id)
		if err != nil {
			message := "No questionnaire with id " + id
			fmt.Println(message)
			panic(message)
		}
		payload, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic("")
		}

		fmt.Println(string(payload))

	},
}

var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "Get the name of a questionnaire",
	Long:  "Get the name of a questionnaire",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		resp, err := http.Get("https://pogues-api.kube.developpement.insee.fr/api/persistence/questionnaire/" + id)
		if err != nil {
			message := "No questionnaire with id " + id
			fmt.Println(message)
			panic(message)
		}
		payload, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic("")
		}
		var questionnaire Questionnaire
		json.Unmarshal(payload, &questionnaire)
		fmt.Println(questionnaire.Name)

	},
}

func init() {
	rootCmd.AddCommand(getCmd)
	getCmd.AddCommand(nameCmd)
}
