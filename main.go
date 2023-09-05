package main

import (
	"fmt"
	"os"
)

func main() {
	var projectName string

	fmt.Print("Enter new project name: ")
	_, err := fmt.Scan(&projectName)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = os.Mkdir(projectName, 0755)

	if err != nil {
        	fmt.Println("Error:", err)
       	 	return
    	}
	
	fmt.Println("Creating project", projectName)
}

