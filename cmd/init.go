/*
Copyright © 2023 Prajwal Raju P <prajwalrajup@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/prajwalraju/RepoReady/impl"
	"github.com/prajwalraju/RepoReady/utils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command that takes the user through the process of creating a new project
var initCmd = &cobra.Command{
	Use:   "init [<folderName> | -n <folderName>] ",
	Short: "Initialize a new project",
	Long:  `Initialize a new project with the given name. create a new folder with the given name and initialize git in it`,
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		var folderName string
		var err error

		// Check if folder name is provided
		folderName, err = cmd.Flags().GetString("folderName")

		if err != nil {
			fmt.Println("Error in getting folderName flag:", err)
			return
		}

		if folderName == "" && len(args) == 0 {
			folderName, err = utils.TakeInput("Folder name", true, "")
			if err != nil {
				fmt.Println("Error in taking folder name:", err)
				return
			}
		} else if folderName == "" {
			folderName = args[0]
		}

		impl.Init(folderName)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().StringP("folderName", "n", "", "Specify the Name of folder or Project")
}
