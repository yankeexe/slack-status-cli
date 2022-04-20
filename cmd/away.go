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
	Long:  "Set your status to away with DND",
	Run: func(cmd *cobra.Command, args []string) {
		client, err := client.GetClient()
		utils.CheckIfError(err)
		client.SetUserPresence("away")

		dnd, err := cmd.Flags().GetInt("dnd")
		utils.CheckIfError(err)

		if dnd != 0 {
			_, err := client.SetSnooze(dnd)
			utils.CheckIfError(err)
			fmt.Println(fmt.Sprintf("🚫 DND enabled for %d minutes", dnd))
		}
		fmt.Println("✅ Status set to away!")
	},
}

func init() {
	rootCmd.AddCommand(awayCmd)
	awayCmd.Flags().Int("dnd", 0, "set do not disturb mode with duration in minutes")
}
