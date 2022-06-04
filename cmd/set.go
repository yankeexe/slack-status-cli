/*
Copyright © 2022 YANKEE MAHARJAN

*/
package cmd

import (
	"log"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/yankeexe/slack-status-cli/internal/config"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set status for slack profile",
	Long:  `Set status for slack profile`,
	Run: func(cmd *cobra.Command, args []string) {
		c := config.Config{}
		c.Load()
		/*
			set shows the list of statuses that are stored in the default profile and the global profile.
		*/
		status := ""
		prompt := &survey.Select{
			Message: "Choose a status:",
			Options: c.GetProfileStatus(),
		}
		survey.AskOne(prompt, &status)
		statusDetails := c.Profiles[c.Default.Name].StatusList[status]
		log.Printf("Status Details: %+v\n", statusDetails)

		// err := client.SetUserCustomStatus()
		// utils.CheckIfError(err)
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
