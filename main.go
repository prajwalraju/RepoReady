package main

import (
	"fmt"
	"repoReady/utils"
)

func main() {
	var projectName string

	projectName, err := utils.TakeInputFromUser("Enter new project name: ", true)
	if err != nil {
		fmt.Println("Error in taking input:", err)
		return
	}

	_, err = utils.CreateFolder(projectName)
	if err != nil {
		fmt.Println("Error in creating folder:", err)
		return
	}

}
