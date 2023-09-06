/*
Copyright © 2023 Prajwal Raju P <prajwalrajup@gmail.com>
*/
package cmd

import (
	"fmt"
	"github.com/prajwalraju/RepoReady/utils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project",
	Long:  `Initialize a new project with the given name. create a new folder with the given name and initialize git in it`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) == 0 {
			cmd.Usage()
			return
		}

		folderName := args[0]
		_, err := utils.CreateFolder(folderName)
		if err != nil {
			fmt.Println("Error in creating folder:", err)
		}

		err = utils.RunGitInit(folderName)
		if err != nil {
			fmt.Println("Error in initializing git:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
