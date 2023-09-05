package utils

import (
	"errors"
	"fmt"
	"os"
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

// Take input from user with a additional field to check if the input is empty
func TakeInputFromUser(message string, checkEmpty bool) (string, error) {
	var input string

	fmt.Print(message)
	_, err := fmt.Scan(&input)

	if err != nil {
		return "", err
	}

	if checkEmpty && input == "" {
		return "", errors.New("input cannot be empty")
	}

	return input, nil
}
