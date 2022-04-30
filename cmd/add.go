/*
Copyright © 2022 YANKEE MAHARJAN

*/
package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
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

Options for the duration:
minute: m, min, mins       :: Example: 10 m, 10 mins, 10minute, 10 minutes
hour:   h, hr, hour, hours :: Example: 1 h, 1hr, 1 hour, 1 hours
day:    d, day, days       :: Example: 2d, 2 day, 2 days
`,
	Run: func(cmd *cobra.Command, args []string) {
		c := config.Config{}
		c.Load()
		// statusStore := config.StatusStore{}

		answers := struct {
			Status   string
			Duration string
		}{}

		var qs = []*survey.Question{
			{
				Name:     "Status",
				Prompt:   &survey.Input{Message: "Status"},
				Validate: survey.Required,
			},
			{
				Name:     "Duration",
				Prompt:   &survey.Input{Message: "Message duration", Help: "Example: 25 mins, 2 hours, 1 day | Defaults to minutes"},
				Validate: survey.Required,
			},
		}
		err := survey.Ask(qs, &answers)
		utils.CheckIfError(err)
		utils.ParseDuration(answers.Duration)

		fmt.Println(answers.Status, answers.Duration)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
