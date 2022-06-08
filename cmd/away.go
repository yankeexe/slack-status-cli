/*
Copyright © 2022 YANKEE MAHARJAN
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yankeexe/slack-status-cli/internal/client"
	"github.com/yankeexe/slack-status-cli/internal/utils"
)

// awayCmd represents the away command
var awayCmd = &cobra.Command{
	Use:   "away",
	Short: "Set your status to away with DND",
	Long: `Set your status to away with DND

Valid durations for --dnd flag includes:
- minutes, hours or days.
DEFAULTS to minutes

NOTE: use single or double quotes around the value if contains space.
OPTIONS for the duration:
minute: m, min, mins       :: Example: "10 m", "10 mins", 10minute, "10 minutes"
hour:   h, hr, hour, hours :: Example: "1 h", 1hr, "1 hour", "1 hours"
day:    d, day, days       :: Example: 2d, "2 day", 2days

`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := client.GetClient()
		utils.CheckIfError(err)
		client.SetUserPresence("away")

		dnd, err := cmd.Flags().GetString("dnd")
		utils.CheckIfError(err)

		if len(dnd) != 0 && dnd != "0" {
			parsed := utils.ParseDuration(dnd)
			_, err := client.SetSnooze(parsed.AbsolutePeriod)
			utils.CheckIfError(err)
			fmt.Printf("🚫 DND enabled for %v\n", parsed.UserDefinedDuration)
		}
		fmt.Println("✅ Status set to away!")
	},
}

func init() {
	rootCmd.AddCommand(awayCmd)
	awayCmd.Flags().String("dnd", "", "set do not disturb mode with user-defined duration")
}
