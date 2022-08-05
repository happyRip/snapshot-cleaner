/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/happyRip/snapshot-cleaner/pkg/hub"
)

var rootCmd = &cobra.Command{
	Use:   "snapshot-cleaner",
	Short: "A simple tool to cleanup old docker hub tags.",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

var client *hub.Client

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	var err error
	client, err = hub.NewClient()
	if err != nil {
		log.Fatal(err)
	}
}
