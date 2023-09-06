package utils

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"os/exec"
)

// CreateFolder creates a new folder with the given name
func CreateFolder(fileName string) (bool, error) {

	if fileName == "" || len(fileName) == 0 || fileName == " " {
		return false, errors.New("file name cannot be empty")
	}

	// 0755 is the default permission mode for directories
	fmt.Println("Creating folder", fileName)
	err := os.Mkdir(fileName, 0755)

	if err != nil {
		fmt.Println("Error in creating folder:", err)
		return false, err
	}

	return true, nil
}

func RunGitInit(fileName string) error {
	if fileName == "" || len(fileName) == 0 || fileName == " " {
		return errors.New("file name cannot be empty")
	}
	cmd := exec.Command("git", "init")
	cmd.Dir = fileName
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
