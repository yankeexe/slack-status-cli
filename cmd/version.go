/*
Copyright (c) Yankee Maharjan 2022. All rights reserved.
Licensed under the MIT license. See LICENSE file in the project root for full license information.
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "development"

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Shows the application version",
	Long:    `Show the applicaton version`,
	Example: "st version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:", Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
