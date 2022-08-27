package client

import (
	"errors"

	"github.com/slack-go/slack"
	"github.com/spf13/viper"
)

func GetClient() (*slack.Client, error) {
	if viper.GetString("default.token") == "" {
		return nil, errors.New(`default profile not set
use:
st profile --default
to set default profile
`)
	}
	return slack.New(viper.GetString("default.token")), nil
}
