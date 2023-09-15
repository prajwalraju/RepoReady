package utils

import (
	"os"
	"testing"
)

func TestCreateFolder(t *testing.T) {

	var err error
	folderName := "TestFolderCreateFolder"

	if err := CreateFolder(folderName); err != nil {
		t.Errorf("Error in creating folder: %v", err)
	}

	if _, err = os.Stat(folderName); os.IsNotExist(err) {
		t.Errorf("Folder dose not exists: %v", err)
	}

	if err = os.RemoveAll(folderName); err != nil {
		t.Errorf("Error in removing folder: %v", err)
	}

}

func TestCreateFolderWithEmptyName(t *testing.T) {

	folderName := ""
	if err := CreateFolder(folderName); err == nil {
		t.Errorf("Error in creating folder: %v", err)
	}

}

func TestRunGitInit(t *testing.T) {
	var err error
	folderName := "TestFolderGit"

	if err = CreateFolder(folderName); err != nil {
		t.Errorf("Error in creating folder: %v", err)
	}
	err = RunGitInit(folderName)

	if err != nil {
		t.Errorf("Error in initializing git: %v", err)
	}

	if _, err = os.Stat(folderName); os.IsNotExist(err) {
		t.Errorf("Folder dose not exists: %v", err)
	}

	if _, err = os.Stat(folderName + "/.git"); os.IsNotExist(err) {
		t.Errorf("Folder dose not exists: %v", err)
	}

	if err = os.RemoveAll(folderName); err != nil {
		t.Errorf("Error in removing folder: %v", err)
	}

}

func TestRunGitInitWithEmptyName(t *testing.T) {

	folderName := ""
	if err := RunGitInit(folderName); err == nil {
		t.Errorf("Error in initializing git: %v", err)
	}

}

func TestRunCommand(t *testing.T) {
	var err error
	folderName := "TestFolderCommand"

	if err = CreateFolder(folderName); err != nil {
		t.Errorf("Error in creating folder: %v", err)
	}
	err = RunCommand("git", []string{"init"}, folderName)

	if err != nil {
		t.Errorf("Error in initializing git: %v", err)
	}

	if _, err = os.Stat(folderName); os.IsNotExist(err) {
		t.Errorf("Folder dose not exists: %v", err)
	}

	if _, err = os.Stat(folderName + "/.git"); os.IsNotExist(err) {
		t.Errorf("Folder dose not exists: %v", err)
	}

	if err = os.RemoveAll(folderName); err != nil {
		t.Errorf("Error in removing folder: %v", err)
	}

}

func TestRunCommandWithEmptyName(t *testing.T) {

	folderName := ""
	if err := RunCommand("git", []string{"init"}, folderName); err == nil {
		t.Errorf("Error in initializing git: %v", err)
	}

}

func TestWriteToFile(t *testing.T) {

	fileName := "TestFile"
	content := "Test Content"
	var err error
	if err = WriteToFile(fileName, content); err != nil {
		t.Errorf("Error in writing to file: %v", err)
	}

	if _, err = os.Stat(fileName); os.IsNotExist(err) {
		t.Errorf("File dose not exists: %v", err)
	}

	if err = os.Remove(fileName); err != nil {
		t.Errorf("Error in removing file: %v", err)
	}
}
