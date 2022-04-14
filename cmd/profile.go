/*
Copyright © 2022 YANKEE MAHARJAN
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/yankeexe/slack-status-cli/internal/config"
	"github.com/yankeexe/slack-status-cli/internal/utils"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Create Slack profile for settings status",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.Config{}
		err := viper.Unmarshal(&config)
		utils.CheckIfError(err)
		fmt.Println(config.GetProfiles())
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
