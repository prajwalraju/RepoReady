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

		if utils.GetEnvVar("GITHUB_TOKEN") == "" {
			utils.UserUnauth()
			return
		}

		repoInput := utils.RepoInput{}

		// Create folder with the given name
		folderName := args[0]
		repoInput.Name = folderName

		_, err := utils.CreateFolder(folderName)
		if err != nil {
			fmt.Println("Error in creating folder:", err)
			return
		}

		// Initialize git in the folder
		err = utils.RunGitInit(folderName)
		if err != nil {
			fmt.Println("Error in initializing git:", err)
			return
		}

		// Check if user wants to push the repo to github
		githubPush, err := utils.TakeOptionInput("Push to github with repo name "+folderName, false, []string{"Yes", "No"})

		if err != nil {
			fmt.Println("Error in taking option to push to github input:", err)
			return
		}

		// If yes, create a remote repo
		if githubPush == "Yes" {

			// Option to add description
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

			topicsArr := []string{}
			if topics != "" {
				topicsArr = strings.Split(topics, ",")
			}

			// Option to make repo private
			githubPrivate, err := utils.TakeOptionInput("Make github repo private "+folderName, false, []string{"Yes", "No"})

			if err != nil {
				fmt.Println("Error in taking github private option input:", err)
				return
			}

			// Create remote repo
			repoInput = utils.RepoInput{Name: folderName, Description: desc, Private: githubPrivate == "Yes", Topics: topicsArr}
			result, err := utils.CreateRemoteRepo(repoInput)

			if err != nil {
				fmt.Println("Error in creating remote repo:", err)
				return
			}

			// Update the repoInput with remote repo details
			repoInput.Owner = result.Owner.Login
			repoInput.HtmlUrl = result.HtmlUrl
			repoInput.GitUrl = result.GitUrl
			repoInput.SshUrl = result.SshUrl

			// If topics are provided, update the remote repo with topics
			if len(topicsArr) != 0 {

				// Update the remote repo with topics
				err = utils.UpdateTopics(repoInput)
				if err != nil {
					fmt.Println("Error in updating topics:", err)
				}
			}

			// Add remote
			remoteUrlOption, err := utils.TakeOptionInput("Prefered remote Url option", false, []string{"Https", "SSH"})

			if err != nil {
				fmt.Println("Error in taking remote Url option :", err)
				return
			}
			utils.AddRemote(remoteUrlOption, repoInput)
		}

		// Check if user wants to generate a readme file
		GenerateReadme, err := utils.TakeOptionInput("Generate an Readme file", false, []string{"Yes", "No"})

		if err != nil {
			fmt.Println("Error in taking Readme file generation input:", err)
			return
		}

		if GenerateReadme == "Yes" {
			err = utils.GenerateReadme(folderName, repoInput)
			if err != nil {
				fmt.Println("Error in generating readme:", err)
				return
			}
		}

		// Check if user wants to generate a license file
		GenerateLicense, err := utils.TakeOptionInput("Generate an License file", false, []string{"Yes", "No"})
		if err != nil {
			fmt.Println("Error in taking License file generation input:", err)
			return
		}

		if GenerateLicense == "Yes" {
			if err = utils.GenerateLicense(folderName, ""); err != nil {
				fmt.Println("Error in generating license:", err)
				return
			}
		}

		// Check if user wants to push the repo to github
		if githubPush == "Yes" {
			if err = utils.CommitAndPushToRemote(repoInput); err != nil {
				fmt.Println("Error in committing and pushing to remote:", err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
