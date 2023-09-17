/*
Copyright © 2023 Prajwal Raju P <prajwalrajup@gmail.com>
*/
package cmd

import (
	"fmt"

	"github.com/prajwalraju/RepoReady/dto"
	"github.com/prajwalraju/RepoReady/impl"
	"github.com/spf13/cobra"
)

// licenseCmd represents the license command
var licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Add a license to the project",
	Long:  `This command will generate a license file and add it to the project.`,
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

		impl.GenerateLicense(dto.RepoInput{Name: directory}, licenseType)

		fmt.Println("license called with licenseType:", licenseType)
	},
}

func init() {
	rootCmd.AddCommand(licenseCmd)

	licenseCmd.Flags().StringP("directory", "d", ".", "Directory to add the license file to (default is current directory)")
	licenseCmd.Flags().StringP("licenseType", "l", "", "Specify the license type")
}
