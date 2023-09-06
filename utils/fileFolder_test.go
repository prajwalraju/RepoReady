package utils

import (
	"os"
	"testing"
)

func TestCreateFolder(t *testing.T) {

	folderName := "TestFolderCreateFolder"
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

	err = os.RemoveAll(folderName)
	if err != nil {
		t.Errorf("Error in removing folder: %v", err)
	}

}

func TestCreateFolderWithEmptyName(t *testing.T) {

	folderName := ""
	result, err := CreateFolder(folderName)

	if err == nil {
		t.Errorf("Error in creating folder: %v", err)
	}

	if result {
		t.Errorf("Error in creating folder: %v", err)
	}

}

func TestRunGitInit(t *testing.T) {

	folderName := "TestFolderGit"
	CreateFolder(folderName)
	err := RunGitInit(folderName)

	if err != nil {
		t.Errorf("Error in initializing git: %v", err)
	}

	_, err = os.Stat(folderName)

	if os.IsNotExist(err) {
		t.Errorf("Folder dose not exists: %v", err)
	}

	_, err = os.Stat(folderName + "/.git")

	if os.IsNotExist(err) {
		t.Errorf("Folder dose not exists: %v", err)
	}

	err = os.RemoveAll(folderName)
	if err != nil {
		t.Errorf("Error in removing folder: %v", err)
	}

}

func TestRunGitInitWithEmptyName(t *testing.T) {

	folderName := ""
	err := RunGitInit(folderName)

	if err == nil {
		t.Errorf("Error in initializing git: %v", err)
	}

}
