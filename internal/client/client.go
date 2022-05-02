package client

import (
	"errors"

	"github.com/slack-go/slack"
	"github.com/spf13/viper"
)

func GetClient() (*slack.Client, error) {
	if viper.GetString("default.token") == "" {
		return nil, errors.New("slack token not set")
	}
	return slack.New(viper.GetString("default.token")), nil
}
