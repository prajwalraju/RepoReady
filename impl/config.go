package impl

import (
	"fmt"
	"github.com/prajwalraju/RepoReady/dto"
	"github.com/prajwalraju/RepoReady/utils"
	"gopkg.in/yaml.v3"
)

func InitConfig(config dto.Config) {
	var err error
	homeDir := utils.GetUserHomeDir()
	if homeDir == "" {
		fmt.Println("Error in getting current user:", err)
		return
	}

	fileExists := utils.CheckIfFileOrFolderExists(homeDir + "/.repoready/config.yaml")

	// Create config file if the file dose not exist
	if !fileExists {

		if !utils.CheckIfFileOrFolderExists(homeDir + "/.repoready") {
			if err = utils.CreateFolder(homeDir + "/.repoready"); err != nil {
				fmt.Println("Error in creating folder:", err)
				return
			}
		}

		fmt.Println("Please visit https://github.com/settings/tokens?type=beta to generate a new token.")
		config, err = TakeInput(config)
		if err != nil {
			fmt.Println("Error in taking input:", err)
			return
		}

		// Marshal the data into YAML format
		yamlData, err := yaml.Marshal(&config)
		if err != nil {
			fmt.Println("Error in marshalling data:", err)
			return
		}

		if err = utils.WriteToFileWithByts(homeDir+"/.repoready/config.yaml", yamlData); err != nil {
			fmt.Println("Error in writing to file:", err)
			return
		}
	}

	// Read the config file

}
func TakeInput(repoInput dto.Config) (dto.Config, error) {
	if repoInput.Github.Token == "" {

		githubToken, err := utils.TakeInput("Enter github token", true, "")
		if err != nil {
			fmt.Println("Error in taking github token input:", err)
			return repoInput, err
		}
		repoInput.Github.Token = githubToken
	}

	return repoInput, nil

}
