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
	err := os.Mkdir(fileName, 0755)

	if err != nil {
		fmt.Println("Error in creating folder:", err)
		return false, err
	}

	return true, nil
}

func RunGitInit(fileName string) error {

	return RunCommand("git", []string{"init"}, fileName)
}

func RunCommand(command string, args []string, dir string) error {
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

func GetEnvVar(key string) string {
	return os.Getenv(key)
}

func writeToFile(file *os.File, content string) error {
	if _, err := fmt.Fprintln(file, content); err != nil {
		return err
	}
	return nil
}
