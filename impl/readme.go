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
	content := ""

	config, err := utils.GetEnvVar()
	if err != nil {
		fmt.Println("Error in getting env var:", err)
		return ""
	}

	for topic, topicType := range config.Readme.Template {
		if topicType == "string" {
			concatString(&content, "# "+topic)
			concatString(&content, repoInput.Name)
			concatString(&content, "\n")
		} else if topicType == "list" {
			concatString(&content, "# "+topic)
			for _, l := range repoInput.Features {
				concatString(&content, "- "+l)
			}
			concatString(&content, "\n")
		}
	}
	return content
}

func concatString(content *string, str string) {
	*content = *content + str + "\n"
}

// CollectAllInfo collects all the info required to create a remote repo
func CollectAllInfo(repoInput dto.RepoInput) (dto.RepoInput, error) {

	// Declare variables
	var err error

	// Option to add Repo name
	if repoInput.Name == "" {
		if repoInput.Name, err = utils.TakeInput("Name of repo", false, ""); err != nil {
			fmt.Println("Error in taking description input:", err)
			return repoInput, err
		}
	}

	// Option to add description
	if repoInput.StringMetadata["LongDescription"], err = utils.TakeInput("Longer Description", false, ""); err != nil || repoInput.StringMetadata["LongDescription"] == "" {
		repoInput.Description = repoInput.StringMetadata["LongDescription"]
	}

	// Option to add Features
	if feature, err := utils.TakeInput("Enter list of featurs (Enter empty to stop)", false, ""); err != nil {
		fmt.Println("Error in taking features input:", err)
		return repoInput, err
	} else if feature != "" {
		repoInput.Features = append(repoInput.Features, feature)

	whileLoop:
		for {
			if feature, err = utils.TakeInput("- ", false, ""); err != nil {
				fmt.Println("Error in taking features input:", err)
				return repoInput, err
			} else if feature != "" {
				repoInput.Features = append(repoInput.Features, feature)
			} else {
				break whileLoop
			}
		}
	}

	config, err := utils.GetEnvVar()
	if err != nil {
		fmt.Println("Error in getting env var:", err)
		return repoInput, err
	}

	for topic, topicType := range config.Readme.Template {

		if topicType == "string" {
			if repoInput.Name, err = utils.TakeInput(topic, false, ""); err != nil {
				fmt.Println("Error in taking description input:", err)
				return repoInput, err
			}
		} else if topicType == "list" {
			if feature, err := utils.TakeInput(topic, false, ""); err != nil {
				fmt.Println("Error in taking features input:", err)
				return repoInput, err
			} else if feature != "" {
				repoInput.Features = append(repoInput.Features, feature)

				// whileLoop:
				// 	for {
				// 		if feature, err = utils.TakeInput("- ", false, ""); err != nil {
				// 			fmt.Println("Error in taking features input:", err)
				// 			return repoInput, err
				// 		} else if feature != "" {
				// 			repoInput.Features = append(repoInput.Features, feature)
				// 		} else {
				// 			break whileLoop
				// 		}
				// 	}
				// }
			}
		}
		return repoInput, nil
	}

	return repoInput, nil
}
