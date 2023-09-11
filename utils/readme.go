package utils

import (
	"fmt"
	"os"
)

func GenerateReadme(directory string, repoInput RepoInput) error {

	fmt.Println("Generating readme file")

	filePath := directory + "/README.md"

	// Open the file for writing (create if it doesn't exist)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	if repoInput, err = CollectAllInfo(repoInput); err != nil {
		fmt.Println("Error in taking inputs:", err)
		return err
	}

	// Write the content to the file
	content := BuildReadmeFileContent(repoInput)
	if err = writeToFile(file, content); err != nil {
		fmt.Println("Error in writing to file:", err)
		return err
	}

	return nil
}

func BuildReadmeFileContent(repoInput RepoInput) string {
	// Name
	var content string
	content += "# " + repoInput.Name + "\n"

	return content
}

// CollectAllInfo collects all the info required to create a remote repo
func CollectAllInfo(repoInput RepoInput) (RepoInput, error) {

	// Option to add description
	if repoInput.Name == "" {
		name, err := TakeInput("Name of repo", false, "")

		if err != nil {
			fmt.Println("Error in taking description input:", err)
			return repoInput, err
		}

		repoInput.Name = name
	}

	return repoInput, nil
}
