/*
Copyright © 2023 Prajwal Raju P <prajwalrajup@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "RepoReady",
	Short: "A CLI to streamline GitHub project setup.",
	Long: `Your All-in-One GitHub Project Setup Assistant

RepoReady is a versatile command-line tool designed to simplify the process of creating and configuring GitHub repositories. 
With RepoReady, you can swiftly set up new projects on GitHub, complete with an informative README file and an appropriate LICENSE.
	
Say goodbye to the hassle of manually crafting README files and LICENSE documents. 
RepoReady automates these tasks for you, ensuring that your GitHub projects are not only well-structured but also legally compliant.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
