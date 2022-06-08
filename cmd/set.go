/*
Copyright © 2022 YANKEE MAHARJAN

*/
package cmd

import (
	"log"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/yankeexe/slack-status-cli/internal/client"
	"github.com/yankeexe/slack-status-cli/internal/config"
	"github.com/yankeexe/slack-status-cli/internal/utils"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set status for slack profile",
	Long:  `Set status for slack profile`,
	Run: func(cmd *cobra.Command, args []string) {
		c := config.Config{}
		c.Load()
		client, err := client.GetClient()
		utils.CheckIfError(err)

		away, err := cmd.Flags().GetBool("away")
		utils.CheckIfError(err)
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

		statusPeriod := time.Now().Add(time.Duration(statusDetails.Period) * time.Minute).Unix()
		log.Println("Status Period UNIX", statusPeriod)

		err = client.SetUserCustomStatus(statusDetails.Status, statusDetails.Emoji, statusPeriod)
		utils.CheckIfError(err)

		if away {
			err = client.SetUserPresence("away")
			utils.CheckIfError(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.Flags().BoolP("away", "", false, "Set status to away")

}
