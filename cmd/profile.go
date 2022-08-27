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
		selectProfile, _ := cmd.Flags().GetBool("select")
		showProfile, _ := cmd.Flags().GetBool("show")

		if !manageProfile && !selectProfile && flagCount > 1 {
			color.Red.Println("Cannot process these flags together")
			os.Exit(1)
		}

		if showProfile {
			if len(config.Default.Name) > 0 {
				color.Green.Println("Current profile is:", config.Default.Name)
				os.Exit(0)
			} else {
				color.Red.Println(`Default slack profile not set
use:
st profile -d
to select a default profile`)
				os.Exit(1)
			}

		}

		if createNew {
			handleCreateNewProfile(&config)
		}

		if manageProfile {
			handleManageProfile(&config, selectProfile)
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
	profileCmd.Flags().BoolP("select", "s", false, "select profile for management, use with --manage")
	profileCmd.Flags().Bool("show", false, "show current profile")

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

func handleManageProfile(c *config.Config, selectProfile bool) {
	selectedProfile := ""
	selectedAction := ""
	profiles := c.GetProfiles()

	if !selectProfile && len(c.Default.Name) != 0 {
		selectedProfile = c.Default.Name
	}

	if !selectProfile && len(c.Default.Name) == 0 {
		color.Red.Println(`Default slack profile not set
use:
st profile -d
to select a default profile`)
		os.Exit(1)
	}

	if len(profiles) == 0 {
		utils.HandleNoProfiles()
	}

	if selectProfile {
		prompt := &survey.Select{
			Message: "Select profile to manage:",
			Options: profiles,
		}
		survey.AskOne(prompt, &selectedProfile)
	}

	actionPrompt := &survey.Select{
		Message: fmt.Sprintf("Select action on profile [%s]", selectedProfile),
		Options: []string{"Rename profile", "Update token", "Delete", "Edit status"},
	}
	log.Println("Selected profiles", selectedProfile)
	survey.AskOne(actionPrompt, &selectedAction)
	log.Println("Selected Action", selectedAction)

	switch selectedAction {
	case "Delete":
		fmt.Println("Proceeding to delete")
		c.DeleteProfile(selectedProfile)
	case "Rename profile":
		c.RenameProfile(selectedProfile)

	case "Update token":
		c.UpdateToken(selectedProfile)

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
