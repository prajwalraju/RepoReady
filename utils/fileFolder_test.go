package utils

import (
	"os"
	"strings"
	"testing"
)

func TestCreateFolder(t *testing.T) {

	folderName := "TestFolder"
	result, err := CreateFolder(folderName)

	if err != nil {
		t.Errorf("Error in creating folder: %v", err)
	}

	if !result {
		t.Errorf("Error in creating folder: %v", err)
	}

	_, err = os.Stat(folderName)

	if os.IsNotExist(err) {
		t.Errorf("Folder dose not exists: %v", err)
	}

	os.Remove(folderName)

}

func TestTakeInputFromUser(t *testing.T) {
	SendInput("Test\n", t)
	input, err := TakeInputFromUser("Enter input: ", true)

	if err != nil {
		t.Errorf("Error in taking input: %v", err)
	}

	if input == "" {
		t.Errorf("Input cannot be empty")
	}
}

func SendInput(input string, t *testing.T) {

	rdr := strings.NewReader(input)
}
