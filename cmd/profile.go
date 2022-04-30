/*
Copyright © 2022 YANKEE MAHARJAN
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/yankeexe/slack-status-cli/internal/config"
	"github.com/yankeexe/slack-status-cli/internal/utils"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Create and manage Slack profiles",
	Run: func(cmd *cobra.Command, args []string) {
		config := config.Config{}
		config.Load()

		createNew, err := cmd.Flags().GetBool("create")
		utils.CheckIfError(err)
		manageProfile, err := cmd.Flags().GetBool("manage")
		utils.CheckIfError(err)

		if createNew && manageProfile {
			color.Red.Println("Cannot use both create and manage flag at the same time")
			os.Exit(1)
		}

		if createNew {
			handleCreateNewProfile(&config)
		}

		if manageProfile {
			handleManageProfile()
		}

		profiles := config.GetProfiles()
		if len(profiles) == 0 {
			handleNoProfiles()
		}
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
	profileCmd.Flags().BoolP("create", "c", false, "create a new slack profile")
	profileCmd.Flags().BoolP("manage", "m", false, "manage slack profile")
}

func handleNoProfiles() {
	color.Red.Println(`No profiles found.

$ st profile --create
to create a new slack profile`)

	os.Exit(1)
}

func handleCreateNewProfile(c *config.Config) {
	profileInfo := config.ProfileInfo{}
	var qs = []*survey.Question{
		{
			Name:     "name",
			Prompt:   &survey.Input{Message: "Profile name"},
			Validate: survey.Required,
		},
		{
			Name:     "token",
			Prompt:   &survey.Password{Message: "User OAuth Token"},
			Validate: survey.Required,
		},
	}

	err := survey.Ask(qs, &profileInfo)
	utils.CheckIfError(err)
	log.Println("New profile", profileInfo)
	c.AddProfile(profileInfo)

	os.Exit(0)
}

func handleManageProfile() {
	fmt.Println("Managing profile.")
	os.Exit(0)
}
