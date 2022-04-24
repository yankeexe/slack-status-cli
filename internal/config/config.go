package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gookit/color"
	"github.com/pelletier/go-toml"
	"github.com/spf13/viper"
	"github.com/yankeexe/slack-status-cli/internal/utils"
)

type ProfileInfo struct {
	Name  string
	Token string
}

type Config struct {
	Globals  []string
	Default  string
	Profiles map[string]Profile
}

type Profile struct {
	Token  string   `mapstructure:"token"`
	Status []string `mapstructure:"Status"`
}

func (c *Config) Load() {
	err := viper.Unmarshal(&c)
	utils.CheckIfError(err)
}

func (c *Config) GetProfiles() []string {
	profiles := []string{}
	for k, _ := range c.Profiles {
		profiles = append(profiles, k)
	}
	return profiles
}

func (c *Config) AddProfile(profileInfo ProfileInfo) {
	fmt.Println("Adding profile")
	viper.Unmarshal(&c)
	profiles := c.GetProfiles()

	if len(profiles) != 0 {
		for _, profile := range c.GetProfiles() {
			if profile == profileInfo.Name {
				color.Red.Println(`Profile already exists!
To edit profile information: st profile --edit`)
				os.Exit(1)
			}
		}
		//currentProfiles := viper.GetStringMapString("profiles")
		newProfile := make(map[string]Profile)
		newProfile[profileInfo.Name] = Profile{Token: profileInfo.Token}
		c.Profiles = newProfile
		data, err := toml.Marshal(c)
		utils.CheckIfError(err)
		fmt.Println("Config", c)
		fmt.Println("Writing to config file...")
		ioutil.WriteFile("/home/yankee/.slack-status-cli/config", data, 0777)
	}
}
