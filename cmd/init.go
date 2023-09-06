/*
Copyright © 2023 Prajwal Raju P <prajwalrajup@gmail.com>
*/
package cmd

import (
	"fmt"
	"strings"

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
			return
		}

		err = utils.RunGitInit(folderName)
		if err != nil {
			fmt.Println("Error in initializing git:", err)
			return
		}

		githubPush, err := utils.TakeOptionInput("Push to github with repo name "+folderName, false, []string{"Yes", "No"})

		if err != nil {
			fmt.Println("Error in taking description input:", err)
			return
		}

		if githubPush == "Yes" {

			desc, err := utils.TakeInput("Description", false, "")

			if err != nil {
				fmt.Println("Error in taking description input:", err)
				return
			}

			// Option to add topics
			topics, err := utils.TakeInput("Tags / Topics (comma seperated)", false, "")

			if err != nil {
				fmt.Println("Error in taking topics input:", err)
				return
			}

			topicsArr := strings.Split(topics, ",")

			// Option to make repo private
			githubPrivate, err := utils.TakeOptionInput("Make github repo private "+folderName, false, []string{"Yes", "No"})

			if err != nil {
				fmt.Println("Error in taking description input:", err)
				return
			}

			repoInput := utils.RepoInput{Name: folderName, Description: desc, Private: githubPrivate == "Yes", HasWiki: true, Topics: topicsArr}
			result, err := utils.CreateRemoteRepo(repoInput)

			if err != nil {
				fmt.Println("Error in creating remote repo:", err)
				return
			}

			fmt.Println("Repo url:", result)

		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
