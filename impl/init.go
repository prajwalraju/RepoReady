package impl

import (
	"fmt"
	"strings"

	"github.com/prajwalraju/RepoReady/dto"
	"github.com/prajwalraju/RepoReady/utils"
)

func Init(folderName string) {

	repoInput := dto.RepoInput{}
	var err error

	// Create folder with the given name
	repoInput.Name = folderName

	if err := utils.CreateFolder(folderName); err != nil {
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

		// Read the config file
		config, err := utils.GetEnvVar()

		if err != nil || config.Github.Token == "" {
			utils.UserUnauth()
			return
		}

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
		repoInput = dto.RepoInput{Name: folderName, Description: desc, Private: githubPrivate == "Yes", Topics: topicsArr}
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
	GenerateReadmeFlag, err := utils.TakeOptionInput("Generate an Readme file", false, []string{"Yes", "No"})

	if err != nil {
		fmt.Println("Error in taking Readme file generation input:", err)
		return
	}

	if GenerateReadmeFlag == "Yes" {
		err = GenerateReadme(folderName, repoInput)
		if err != nil {
			fmt.Println("Error in generating readme:", err)
			return
		}
	}

	// Check if user wants to generate a license file
	GenerateLicenseFlag, err := utils.TakeOptionInput("Generate an License file", false, []string{"Yes", "No"})
	if err != nil {
		fmt.Println("Error in taking License file generation input:", err)
		return
	}

	if GenerateLicenseFlag == "Yes" {
		if err = GenerateLicense(folderName, ""); err != nil {
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
}
