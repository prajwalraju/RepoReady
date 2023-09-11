/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/prajwalraju/RepoReady/utils"
	"github.com/spf13/cobra"
)

// licenseCmd represents the license command
var licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		directory, err := cmd.Flags().GetString("directory")
		if err != nil {
			fmt.Println("Error in getting directory flag:", err)
			return
		}

		licenseType, err := cmd.Flags().GetString("licenseType")
		if err != nil {
			fmt.Println("Error in getting licenseType flag:", err)
			return
		}

		utils.GenerateLicense(directory, licenseType)

		fmt.Println("license called with licenseType:", licenseType)
	},
}

func init() {
	rootCmd.AddCommand(licenseCmd)

	licenseCmd.Flags().StringP("directory", "d", ".", "Directory to add the license file to (default is current directory)")
	licenseCmd.Flags().StringP("licenseType", "l", "", "Specify the license type")
}
