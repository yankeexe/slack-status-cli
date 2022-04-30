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
	"github.com/spf13/viper"
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

		createNew, _ := cmd.Flags().GetBool("create")
		manageProfile, _ := cmd.Flags().GetBool("manage")
		defaultProfile, _ := cmd.Flags().GetBool("default")
		flagCount := cmd.Flags().NFlag()

		if flagCount > 1 {
			color.Red.Println("accepts at most 1 flag, received", flagCount)
			os.Exit(1)
		}

		if createNew {
			handleCreateNewProfile(&config)
		}

		if manageProfile {
			handleManageProfile(&config)
		}

		if defaultProfile {
			handleSetDefaultProfile(&config)
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
	profileCmd.Flags().BoolP("default", "d", false, "select default profile")
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

func handleManageProfile(c *config.Config) {
	log.Println("Managing profile.")
	selectedProfile := ""
	prompt := &survey.Select{
		Message: "Select profile to manage:",
		Options: c.GetProfiles(),
	}
	survey.AskOne(prompt, &selectedProfile)
	log.Println("Selected profiles", selectedProfile)
	os.Exit(0)
}

func handleSetDefaultProfile(c *config.Config) {
	log.Println("Setting default profile")
	selectedProfile := ""
	prompt := &survey.Select{
		Message: "Select profile to set as default:",
		Options: c.GetProfiles(),
	}
	survey.AskOne(prompt, &selectedProfile)
	c.Default.Name = selectedProfile
	c.Default.Token = viper.GetString(fmt.Sprintf("profiles.%s.token", selectedProfile))
	log.Println("Selected profiles", selectedProfile)
	c.Save()
}
