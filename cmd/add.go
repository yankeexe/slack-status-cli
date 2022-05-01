/*
Copyright © 2022 YANKEE MAHARJAN

*/
package cmd

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/yankeexe/slack-status-cli/internal/config"
	"github.com/yankeexe/slack-status-cli/internal/utils"
	"log"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add status for default Slack profile",
	Long: `Add status for default Slack profile

Valid durations include:
- minutes, hours or days.

Options for the duration:
minute: m, min, mins       :: Example: 10 m, 10 mins, 10minute, 10 minutes
hour:   h, hr, hour, hours :: Example: 1 h, 1hr, 1 hour, 1 hours
day:    d, day, days       :: Example: 2d, 2 day, 2 days
`,
	Run: func(cmd *cobra.Command, args []string) {
		c := config.Config{}
		c.Load()
		statusContainer := config.StatusStore{}

		var qs = []*survey.Question{
			{
				Name:     "Status",
				Prompt:   &survey.Input{Message: "Status"},
				Validate: survey.Required,
			},
			{
				Name:     "UserDefinedDuration",
				Prompt:   &survey.Input{Message: "Message duration", Help: "Example: 25 mins, 2 hours, 1 day | Defaults to minutes"},
				Validate: survey.Required,
			},
		}
		err := survey.Ask(qs, &statusContainer)
		utils.CheckIfError(err)
		statusContainer.ParseDuration()
		log.Printf("Status Container: %+v\n", statusContainer)

		c.AddStatus(statusContainer)

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

}
