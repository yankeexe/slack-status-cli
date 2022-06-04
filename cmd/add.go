/*
Copyright © 2022 YANKEE MAHARJAN

*/
package cmd

import (
	"log"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"github.com/yankeexe/slack-status-cli/internal/config"
	"github.com/yankeexe/slack-status-cli/internal/utils"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add status for default Slack profile",
	Long: `Add status for default Slack profile

Valid durations include:
- minutes, hours or days.
Defaults to minutes

Options for the duration:
minute: m, min, mins       :: Example: 10 m, 10 mins, 10minute, 10 minutes
hour:   h, hr, hour, hours :: Example: 1 h, 1hr, 1 hour, 1 hours
day:    d, day, days       :: Example: 2d, 2 day, 2 days
`,
	Run: func(cmd *cobra.Command, args []string) {
		global, err := cmd.Flags().GetBool("global")
		utils.CheckIfError(err)
		c := config.Config{}
		c.Load()
		if len(c.Default.Name) == 0 {
			color.Red.Println("No Slack profiles found to add status.")
			color.Green.Println(`Create a slack profile using:
$ st profile --create`)
			os.Exit(1)
		}
		statusContainer := config.StatusStore{}

		var qs = []*survey.Question{
			{
				Name:     "Status",
				Prompt:   &survey.Input{Message: "Status"},
				Validate: survey.Required,
			},
			{
				Name:     "UserDefinedDuration",
				Prompt:   &survey.Input{Message: "Duration", Help: "Example: 25 mins, 2 hours, 1 day | Defaults to minutes"},
				Validate: survey.Required,
			},
			{Name: "Emoji",
				Prompt: &survey.Select{
					Message: "Choose an emoji:",
					Options: utils.EmojiKeys(),
				},
				Validate: survey.Required,
			},
		}
		err = survey.Ask(qs, &statusContainer)
		utils.CheckIfError(err)

		parsed := utils.ParseDuration(statusContainer.UserDefinedDuration)
		statusContainer.UserDefinedDuration = parsed.UserDefinedDuration
		statusContainer.Period = parsed.AbsolutePeriod
		log.Printf("Status Container: %+v\n", statusContainer)
		if global {
			c.AddGlobalStatus(statusContainer)
			c.Save()
		} else {
			c.AddStatus(statusContainer)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolP("global", "g", false, "add to the global list of statuses")
}
