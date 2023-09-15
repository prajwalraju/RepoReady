/*
Copyright © 2023 Prajwal Raju P <prajwalrajup@gmail.com>
*/
package cmd

import (
	"github.com/prajwalraju/RepoReady/impl"
	"github.com/spf13/cobra"
)

// initCmd represents the init command that takes the user through the process of creating a new project
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new project",
	Long:  `Initialize a new project with the given name. create a new folder with the given name and initialize git in it`,
	Run: func(cmd *cobra.Command, args []string) {

		// Check if folder name is provided
		if len(args) == 0 {
			cmd.Usage()
			return
		}

		impl.Init(args[0])
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
