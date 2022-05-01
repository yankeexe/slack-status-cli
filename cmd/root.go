/*
Copyright © 2022 YANKEE MAHARJAN

*/
package cmd

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "st",
	Short: "Update Slack status with speed",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	usr, _ := user.Current()
	configDirPath := filepath.Join(usr.HomeDir, ".slack-status-cli")
	viper.AddConfigPath(configDirPath)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		//fmt.Fprintln(os.Stderr, viper.ConfigFileUsed())
	}

}
