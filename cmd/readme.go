/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/prajwalraju/RepoReady/utils"
	"github.com/spf13/cobra"
)

// readmeCmd represents the readme command
var readmeCmd = &cobra.Command{
	Use:   "readme",
	Short: "Add a readme to the project",
	Long:  `Readme file generator. Creates a readme file using a template and adds it to the project.`,
	Run: func(cmd *cobra.Command, args []string) {
		directory, err := cmd.Flags().GetString("directory")
		if err != nil {
			fmt.Println("Error in getting directory flag:", err)
			return
		}

		err = utils.GenerateReadme(directory)
		if err != nil {
			fmt.Println("Error in generating readme:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(readmeCmd)

	readmeCmd.Flags().StringP("directory", "d", "", "Directory to add the readme to (default is current directory)")
}
