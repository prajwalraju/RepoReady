/*
Copyright © 2023 Prajwal Raju P <prajwalrajup@gmail.com>
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
	Long: `Readme file generator. 
Generats a readme file using a template and adds it to the project.`,
	Run: func(cmd *cobra.Command, args []string) {

		directory, err := cmd.Flags().GetString("directory")
		if err != nil {
			fmt.Println("Error in getting directory flag:", err)
			return
		}

		generateFile, err := cmd.Flags().GetBool("generateFile")
		if err != nil {
			fmt.Println("Error in getting generateFile flag:", err)
			return
		}

		// If generateFile is false, just print the readme file content
		if !generateFile {

			repoInput, err := utils.CollectAllInfo(utils.RepoInput{})
			if err != nil {
				fmt.Println("Error in collecting all info:", err)
				return
			}

			fmt.Println("Readme file content:\n\n")

			fmt.Println(utils.BuildReadmeFileContent(repoInput))
			return
		}

		// Generate readme file
		err = utils.GenerateReadme(directory, utils.RepoInput{})
		if err != nil {
			fmt.Println("Error in generating readme:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(readmeCmd)

	readmeCmd.Flags().StringP("directory", "d", ".", "Directory to add the readme to (default is current directory)")
	readmeCmd.Flags().BoolP("generateFile", "g", true, "Generate a readme file or just print the readme content")
}
