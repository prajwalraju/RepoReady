package impl

import (
	"fmt"
	"strings"

	"github.com/prajwalraju/RepoReady/dto"
	"github.com/prajwalraju/RepoReady/utils"
)

func Init(folderName string) {

	// Create a repoInput object add the folder name to it
	repoInput := dto.RepoInput{Name: folderName}
	var err error

	// Create folder and initialize git in it
	if err = createFolderAndInitRepo(folderName); err != nil {
		return
	}

	// Check if user wants to push the repo to github
	if githubPush, err := utils.TakeOptionInput("Push to github with repo name "+folderName, false, []string{"Yes", "No"}); err != nil {
		fmt.Println("Error in taking option to push to github input:", err)
		return
	} else {
		repoInput.PushToGit = githubPush == "Yes"
	}

	// Check if user wants to push the repo to github
	if repoInput.PushToGit {

		//////////////////////////
		// Get details related to github from the user
		if repoInput, err = getGitHubRelatedDetails(repoInput); err != nil {
			return
		}

		//////////////////////////
		// Create remote repo
		if repoInput, err = createRemoteRepo(repoInput); err != nil {
			return
		}
	}

	// Generate License file in repo if user wants to
	if repoInput, err = generateLicense(folderName, repoInput); err != nil {
		return
	}

	// Generate readme file in repo if user wants to
	if err = generateReadme(folderName, repoInput); err != nil {
		return
	}

	// Check if user wants to push the repo to github
	if repoInput.PushToGit {
		if err = utils.CommitAndPushToRemote(repoInput); err != nil {
			fmt.Println("Error in committing and pushing to remote:", err)
			return
		}
	}
}

// createFolderAndInitRepo creates a folder with the given name and initializes git in it
func createFolderAndInitRepo(folderName string) error {

	// Create folder with the given name
	if err := utils.CreateFolder(folderName); err != nil {
		fmt.Println("Error in creating folder:", err)
		return err
	}

	// Initialize git in the folder
	err := utils.RunGitInit(folderName)
	if err != nil {
		fmt.Println("Error in initializing git:", err)
		return err
	}

	// Return nil if no error
	return nil
}

// GenerateReadme generates a readme file in the given directory
func generateReadme(folderName string, repoInput dto.RepoInput) error {
	// Check if user wants to generate a readme file
	GenerateReadmeFlag, err := utils.TakeOptionInput("Generate an Readme file", false, []string{"Yes", "No"})

	if err != nil {
		fmt.Println("Error in taking Readme file generation input:", err)
		return err
	}

	if GenerateReadmeFlag == "Yes" {
		if err = GenerateReadme(folderName, repoInput); err != nil {
			fmt.Println("Error in generating readme:", err)
			return err
		}
	}

	return nil
}

// GenerateLicense generates a license file in the given directory
func generateLicense(folderName string, repoInput dto.RepoInput) (dto.RepoInput, error) {

	// Check if user wants to generate a license file
	GenerateLicenseFlag, err := utils.TakeOptionInput("Generate an License file", false, []string{"Yes", "No"})
	if err != nil {
		fmt.Println("Error in taking License file generation input:", err)
		return repoInput, err
	}

	if GenerateLicenseFlag == "Yes" {
		if repoInput, err = GenerateLicense(repoInput, ""); err != nil {
			fmt.Println("Error in generating license:", err)
			return repoInput, err
		}
	}

	return repoInput, nil
}

// createRemoteRepo creates a remote repo in github
func createRemoteRepo(repoInput dto.RepoInput) (dto.RepoInput, error) {

	// Declare variables
	var err error

	//////////////////////////
	// Read the config file and check if the user is authenticated
	if config, err := utils.GetEnvVar(); err != nil || config.Github.Token == "" {
		utils.UserUnauth()
		return repoInput, err
	}

	//////////////////////////
	// Create remote repo
	if result, err := utils.CreateRemoteRepo(repoInput); err != nil {
		fmt.Println("Error in creating remote repo:", err)
		return repoInput, err
	} else {
		// Update the repoInput with remote repo details
		repoInput.Owner = result.Owner.Login
		repoInput.HtmlUrl = result.HtmlUrl
		repoInput.GitUrl = result.GitUrl
		repoInput.SshUrl = result.SshUrl
	}

	//////////////////////////
	// If topics are provided, update the remote repo with topics
	if len(repoInput.Topics) != 0 {
		if err = utils.UpdateTopics(repoInput); err != nil {
			fmt.Println("Error in updating topics:", err)
		}
	}

	//////////////////////////
	// Add remote
	if remoteUrlOption, err := utils.TakeOptionInput("Prefered remote Url option", false, []string{"Https", "SSH"}); err != nil {
		fmt.Println("Error in taking remote Url option :", err)
		return repoInput, err
	} else {
		utils.AddRemote(remoteUrlOption, repoInput)
	}

	return repoInput, nil
}

// getGitHubRelatedDetails gets the details related to github from the user
func getGitHubRelatedDetails(repoInput dto.RepoInput) (dto.RepoInput, error) {

	var err error

	//////////////////////////
	// Option to add description
	if repoInput.Description, err = utils.TakeInput("Description", false, ""); err != nil {
		fmt.Println("Error in taking description input:", err)
		return repoInput, err
	}

	//////////////////////////
	// Option to add topics
	if topics, err := utils.TakeInput("Tags / Topics (comma seperated)", false, ""); err != nil {
		fmt.Println("Error in taking topics input:", err)
		return repoInput, err
	} else if topics != "" {
		repoInput.Topics = strings.Split(topics, ",")
	}

	//////////////////////////
	// Option to make repo private
	if githubPrivate, err := utils.TakeOptionInput("Make github repo private "+repoInput.Name, false, []string{"Yes", "No"}); err != nil {
		fmt.Println("Error in taking github private option input:", err)
		return repoInput, err
	} else {
		repoInput.Private = githubPrivate == "Yes"
	}

	return repoInput, nil
}
