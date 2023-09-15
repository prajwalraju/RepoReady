package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"

	"github.com/manifoldco/promptui"
	"github.com/prajwalraju/RepoReady/dto"
	"gopkg.in/yaml.v3"
)

// CreateFolder creates a new folder with the given name
func CreateFolder(fileName string) error {

	if fileName == "" || len(fileName) == 0 || fileName == " " {
		return errors.New("file name cannot be empty")
	}

	// 0755 is the default permission mode for directories
	err := os.Mkdir(fileName, 0755)

	if err != nil {
		fmt.Println("Error in creating folder:", err)
		return err
	}

	return nil
}

func RunGitInit(fileName string) error {

	return RunCommand("git", []string{"init"}, fileName)
}

func RunCommand(command string, args []string, dir string) error {
	if dir == "" {
		return errors.New("dir cannot be empty")
	}

	if command == "" || len(command) == 0 || command == " " {
		return errors.New("command cannot be empty")
	}

	cmd := exec.Command(command, args...)
	if dir != "" {
		cmd.Dir = dir
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func TakeInput(fieldName string, checkIfEmpty bool, def string) (string, error) {
	validate := func(input string) error {
		if len(input) == 0 && checkIfEmpty {
			return errors.New(fieldName + " cannot be empty")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    fieldName,
		Validate: validate,
		Default:  def,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil
}

func TakeOptionInput(fieldName string, checkIfEmpty bool, options []string) (string, error) {
	prompt := promptui.Select{
		Label: fieldName,
		Items: options,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return "", err
	}

	return result, nil

}

func GetEnvVar() (dto.Config, error) {

	file, err := os.Open(GetUserHomeDir() + "/.repoready/config.yaml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return dto.Config{}, err
	}
	defer file.Close()

	// Create a decoder to read the YAML data
	decoder := yaml.NewDecoder(file)

	// Initialize a struct to store the YAML data
	var config dto.Config

	// Decode the YAML data into the struct
	if err := decoder.Decode(&config); err != nil {
		fmt.Println("Failed to decode YAML: ", err)
		return dto.Config{}, err
	}
	return config, nil
}

func WriteToFile(filePath string, content string) error {
	// Open the file for writing (create if it doesn't exist)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	if _, err := fmt.Fprintln(file, content); err != nil {
		return err
	}
	return nil
}

func WriteToFileWithByts(filePath string, content []byte) error {
	// Open the file for writing (create if it doesn't exist)
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer file.Close()

	if _, err := file.Write(content); err != nil {
		return err
	}
	return nil
}

func CheckIfFileOrFolderExists(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}

func GetUserHomeDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error in getting user home directory:", err)
		return ""
	}
	return home
}
