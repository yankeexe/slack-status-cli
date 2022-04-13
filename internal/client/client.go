package client

import (
	"github.com/slack-go/slack"
	"github.com/spf13/viper"
)

func GetClient() *slack.Client {
	return slack.New(viper.GetString("token"))
}
