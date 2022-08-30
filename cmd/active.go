/*
Copyright (c) Yankee Maharjan 2022. All rights reserved.
Licensed under the MIT license. See LICENSE file in the project root for full license information.
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
Removes DND and status if present.`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := client.GetClient()
		utils.CheckIfError(err)
		err = client.SetUserPresence("auto")
		utils.CheckIfError(err)
		err = client.UnsetUserCustomStatus()
		utils.CheckIfError(err)
		err = client.EndDND()
		utils.CheckIfError(err)
		fmt.Println("✅ Status set to active")
	},
}

func init() {
	rootCmd.AddCommand(activeCmd)
}
