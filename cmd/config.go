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

// configCmd represents the setup command

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Setup the configurations and settings for RepoReady",
	Long:  `Setup the configurations, tokens and costomized settings for RepoReady.`,
	Run: func(cmd *cobra.Command, args []string) {
		directory, err := cmd.Flags().GetString("githubToken")
		if err != nil {
			fmt.Println("Error in getting githubToken flag:", err)
			return
		}

		config := dto.Config{Github: dto.Github{Token: directory}}

		impl.InitConfig(config)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringP("githubToken", "g", "", "Github token to be used for creating repositories")
}
