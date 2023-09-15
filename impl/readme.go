package impl

import (
	"fmt"
	"github.com/prajwalraju/RepoReady/dto"
	"github.com/prajwalraju/RepoReady/utils"
)

func GenerateReadme(directory string, repoInput dto.RepoInput) error {

	fmt.Println("Generating readme file")

	filePath := directory + "/README.md"

	var err error
	if repoInput, err = CollectAllInfo(repoInput); err != nil {
		fmt.Println("Error in taking inputs:", err)
		return err
	}

	// Write the content to the file
	content := BuildReadmeFileContent(repoInput)
	if err = utils.WriteToFile(filePath, content); err != nil {
		fmt.Println("Error in writing to file:", err)
		return err
	}

	return nil
}

func BuildReadmeFileContent(repoInput dto.RepoInput) string {
	// Name
	var content string
	content += "# " + repoInput.Name + "\n"

	return content
}

// CollectAllInfo collects all the info required to create a remote repo
func CollectAllInfo(repoInput dto.RepoInput) (dto.RepoInput, error) {

	// Option to add description
	if repoInput.Name == "" {
		name, err := utils.TakeInput("Name of repo", false, "")

		if err != nil {
			fmt.Println("Error in taking description input:", err)
			return repoInput, err
		}

		repoInput.Name = name
	}

	return repoInput, nil
}
