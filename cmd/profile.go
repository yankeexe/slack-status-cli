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
		flagCount := cmd.Flags().NFlag()

		if flagCount == 0 {
			cmd.Help()
			os.Exit(0)
		}

		config := config.Config{}
		config.Load()

		createNew, _ := cmd.Flags().GetBool("create")
		manageProfile, _ := cmd.Flags().GetBool("manage")
		defaultProfile, _ := cmd.Flags().GetBool("default")

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
			utils.HandleNoProfiles()
		}
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
	profileCmd.Flags().BoolP("create", "c", false, "create a new slack profile; sets default profile if not present")
	profileCmd.Flags().BoolP("manage", "m", false, "manage slack profile")
	profileCmd.Flags().BoolP("default", "d", false, "select default profile")
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
	selectedAction := ""
	profiles := c.GetProfiles()

	if len(profiles) == 0 {
		utils.HandleNoProfiles()
	}

	prompt := &survey.Select{
		Message: "Select profile to manage:",
		Options: profiles,
	}
	survey.AskOne(prompt, &selectedProfile)

	actionPrompt := &survey.Select{
		Message: "Select action on profile",
		Options: []string{"Edit Name", "Update token", "Delete", "Edit status"},
	}
	log.Println("Selected profiles", selectedProfile)
	survey.AskOne(actionPrompt, &selectedAction)
	log.Println("Selected Action", selectedAction)

	switch selectedAction {
	case "Delete":
		fmt.Println("Proceeding to delete")
		c.DeleteProfile(selectedProfile)
	}

	os.Exit(0)
}

func handleSetDefaultProfile(c *config.Config) {
	log.Println("Setting default profile")
	selectedProfile := ""
	profiles := c.GetProfiles()

	if len(profiles) == 0 {
		utils.HandleNoProfiles()
	}

	prompt := &survey.Select{
		Message: "Select profile to set as default:",
		Options: profiles,
	}
	survey.AskOne(prompt, &selectedProfile)
	c.Default = config.ProfileInfo{Name: selectedProfile, Token: viper.GetString(fmt.Sprintf("profiles.%s.token", selectedProfile))}
	c.Save()
	color.Green.Println("Default profile:", selectedProfile)
}
