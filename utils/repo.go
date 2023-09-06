package utils

import "fmt"

func AddRemote(remoteUrlOption string, repoInput RepoInput) error {

	commandArgs := []string{"remote", "add", "origin"}
	if remoteUrlOption == "Https" {
		commandArgs = append(commandArgs, repoInput.GitUrl)
	} else {
		commandArgs = append(commandArgs, repoInput.SshUrl)
	}

	if err := RunCommand("git", commandArgs, repoInput.Name); err != nil {
		fmt.Println("Error in adding remote:", err)
		return err
	}
	return nil
}

func CommitAndPushToRemote(repoInput RepoInput) error {
	if err := RunCommand("git", []string{"add", "-A"}, repoInput.Name); err != nil {
		fmt.Println("Error in adding files to git:", err)
		return err
	}

	fmt.Println("Enter commit message: ")
	if err := RunCommand("git", []string{"commit", "-m", "first commit"}, repoInput.Name); err != nil {
		fmt.Println("Error in commiting files:", err)
		return err
	}

	fmt.Println("Enter push message: ")
	if err := RunCommand("git", []string{"push", "-u", "origin", "main"}, repoInput.Name); err != nil {
		fmt.Println("Error in pushing to remote:", err)
		return err
	}
	return nil
}
