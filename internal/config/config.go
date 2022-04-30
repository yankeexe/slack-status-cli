package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"

	"path/filepath"

	"github.com/spf13/viper"

	"github.com/gookit/color"
	"github.com/pelletier/go-toml"
	"github.com/yankeexe/slack-status-cli/internal/utils"
)

var usr, _ = user.Current()

var ConfigDirPath = filepath.Join(usr.HomeDir, ".slack-status-cli")
var ConfigFilePath = filepath.Join(ConfigDirPath, "config")

type Config struct {
	Globals  []string
	Default  ProfileInfo
	Profiles map[string]Profile
}

type ProfileInfo struct {
	Name  string `mapstructure:"name"`
	Token string `mapstructure:"token"`
}

type Profile struct {
	Token  string   `mapstructure:"token"`
	Status []string `mapstructure:"Status"`
}

// Load config file values
func (c *Config) Load() {
	c.exists()
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

/*
	@NOTE:
	Just return if the config directory exists.
	If it does not create one.
	If it exists, do not care if config file is there.
	Contents of config file will be checked later on:
		either nil or has some.
*/
func (c *Config) exists() bool {
	if _, err := os.Stat(ConfigDirPath); os.IsNotExist(err) {
		color.Red.Println("Config directory not found")
		if err := os.Mkdir(ConfigDirPath, 0770); err != nil {
			utils.CheckIfError(err)
			color.Green.Println("Config directory created:", ConfigDirPath)
		}
		return false
	}
	return true
}

func (c *Config) AddProfile(profileInfo ProfileInfo) {
	log.Println("Adding profile")
	profiles := c.GetProfiles()

	if len(profiles) != 0 {
		for _, profile := range c.GetProfiles() {
			if profile == profileInfo.Name {
				color.Red.Println(`Profile already exists!
To edit profile information: st profile --edit`)
				os.Exit(1)
			}
		}
	}
	c.Profiles[profileInfo.Name] = Profile{Token: profileInfo.Token}
	if c.Default.Name == "" {
		c.Default = profileInfo
	}
	c.Save()
	color.Green.Println(fmt.Sprintf("Successfully added Profile: '%s' to config.", profileInfo.Name))
}

func (c *Config) Save() {
	data, err := toml.Marshal(c)
	utils.CheckIfError(err)
	err = ioutil.WriteFile(ConfigFilePath, data, 0777)
	utils.CheckIfError(err)
}
