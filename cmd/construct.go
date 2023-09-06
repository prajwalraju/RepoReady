package cmd

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

/* WARNING - THIS WILL WRITE THE README IN THE DIRECTORY RUNNED
 * PLEASE HAVE CAUTION
 */

// Structure of the README

type PromptStructure struct {
	ErrorMsg string
	Label    string
}

var generation = &cobra.Command{
	Use:   "readme_generate",
	Short: "Will generate and write your README",
	Long:  "Will generate and write your README",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("WARNING - THIS WILL WRITE THE README IN THE DIRECTORY RUNNED PLEASE HAVE CAUTION\n")
		fmt.Printf("\n")
		fmt.Printf("All options are optional, if a option is not provided the value will be 'PlaceHolder'\n")
		fmt.Printf("\n")
		GenerateReadme()
	},
}

func GetUserInput(pc PromptStructure) string {
	// Validate
	validate := func(input string) error {
		return nil
	}

	//Template
	template := &promptui.PromptTemplates{
		Prompt:  "{{ . }}",
		Valid:   "{{ . | green }}",
		Invalid: "{{ . | red }}",
		Success: "{{ . | bold }}",
	}

	// Prompt itself
	prompt := promptui.Prompt{
		Label:     pc.Label,
		Templates: template,
		Validate:  validate,
	}
	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Got the following error: %v\n", err)
		os.Exit(1)
	} else if result == "" {
		fmt.Printf("No value given, using the word PlaceHolder.\n")
		return "PlaceHolder"
	}
	fmt.Printf("Got the following input: %s\n", result)
	return result
}

func GenerateReadme() {

	/* Get user input for the Project Title, Description,
	 * Features and installation process.
	 */

	ProjectTitlePrompt := PromptStructure{
		"Please provide your project title",
		"What is your project title? ",
	}
	ProjectTitle := GetUserInput(ProjectTitlePrompt)

	ProjectDescriptionPrompt := PromptStructure{
		"Please provide your project title",
		"What will be your project description? ",
	}
	ProjectDescription := GetUserInput(ProjectDescriptionPrompt)

	FeaturesPrompt := PromptStructure{
		"Please provide your project features",
		"What is your project features? ",
	}
	Features := GetUserInput(FeaturesPrompt)

	InstallationPrompt := PromptStructure{
		"Please provide how your project installation works.",
		"How your project installation works? ",
	}
	Installation := GetUserInput(InstallationPrompt)

	WriteToReadme(ProjectTitle, ProjectDescription, Features, Installation)
}

func WriteToReadme(ProjectTitle string, ProjectDescription string, Features string, Installation string) {
	// Open the README and create it if needed.
	file, err := os.OpenFile("README.md", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	// Error handling
	if err != nil {
		fmt.Printf("Got the following error: %s\n", err)
	}

	// Write to the readme.
	fmt.Fprintf(file, "# %s\n", ProjectTitle)
	fmt.Fprintf(file, "Welcome to our project, we hope you enjoy your project!\n")
	fmt.Fprintf(file, "# Project description\n")
	fmt.Fprintf(file, "%s\n", ProjectDescription)
	fmt.Fprintf(file, "# Project features\n")
	fmt.Fprintf(file, "Our project features are the following:\n")
	fmt.Fprintf(file, "%s\n", Features)
	fmt.Fprintf(file, "# Installation process\n")
	fmt.Fprintf(file, "%s\n", Installation)
	fmt.Fprintf(file, "# Credits\n")
	fmt.Fprintf(file, "README generator done by aikooo7 for the project RepoReady.\n")
	fmt.Fprintf(file, "https://github.com/prajwalraju/RepoReady\n")
	fmt.Fprintf(file, "Hope you enjoyed it!\n")
	// Close the handler to the readme.
	defer file.Close()
}

func init() {
	rootCmd.AddCommand(generation)
}
