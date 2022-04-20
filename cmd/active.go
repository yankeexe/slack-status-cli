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

// activeCmd represents the active command
var activeCmd = &cobra.Command{
	Use:   "active",
	Short: "Set yourself to active",
	Long: `Set yourself to active.
Removes DND if present.`,
	Run: func(cmd *cobra.Command, args []string) {
		client := client.GetClient()
		err := client.SetUserPresence("auto")
		utils.CheckIfError(err)

		fmt.Println("✅ Status set to active")
		err = client.EndDND()
		utils.CheckIfError(err)
	},
}

func init() {
	rootCmd.AddCommand(activeCmd)
}
