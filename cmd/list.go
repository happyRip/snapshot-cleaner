/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		tags, err := client.ListTags()
		if err != nil {
			log.Fatal(err)
		}
		j, err := json.MarshalIndent(tags, "", "    ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(j))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
