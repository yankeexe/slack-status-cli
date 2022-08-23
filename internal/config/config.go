package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"gopkg.in/yaml.v3"

	"github.com/spf13/viper"

	"github.com/gookit/color"
	"github.com/yankeexe/slack-status-cli/internal/utils"
)

var usr, _ = user.Current()

var ConfigDirPath = filepath.Join(usr.HomeDir, ".slack-status-cli")
var ConfigFilePath = filepath.Join(ConfigDirPath, "config")

type Config struct {
	Default  ProfileInfo
	Globals  []StatusStore
	Profiles map[string]Profile
}

type ProfileInfo struct {
	Name  string `mapstructure:"name"`
	Token string `mapstructure:"token"`
}

type Profile struct {
	Token      string `mapstructure:"token"`
	StatusList map[string]StatusStore
}

// Mapping key: emoji status UserDefinedDuration
type StatusStore struct {
	Status              string
	UserDefinedDuration string
	Period              int
	Emoji               string
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

func (c *Config) GetProfileStatus() []string {
	statusList := []string{}
	profile := c.Profiles[c.Default.Name]
	for k, _ := range profile.StatusList {
		statusList = append(statusList, k)
	}

	log.Printf("Value in slice %+v\n", statusList)
	if len(statusList) == 0 {
		color.Red.Println(`No status found. Please add status using:
$ st add`)
		os.Exit(0)
	}
	return statusList
}

/*
	@NOTE:
	Just return if the config directory exists.
	If it does not create one.
	If it exists, do not care if config file is there.
	Contents of config file will be checked later on:
		either nil or has some.
*/
func (c *Config) exists() {
	if _, err := os.Stat(ConfigDirPath); os.IsNotExist(err) {
		color.Red.Println("Config directory does not exist")
		if err := os.Mkdir(ConfigDirPath, 0770); err != nil {
			utils.CheckIfError(err)
		}
		color.Green.Println("Config directory created:", ConfigDirPath)
	}
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

	// handle nil map
	if c.Profiles == nil {
		c.Profiles = make(map[string]Profile)
	}
	c.Profiles[profileInfo.Name] = Profile{Token: profileInfo.Token}
	if c.Default.Name == "" {
		c.Default = profileInfo
	}
	c.Save()
	color.Green.Println(fmt.Sprintf("Successfully added Profile: '%s' to config.", profileInfo.Name))
}

// AddStatus adds status to the default profile
func (c *Config) AddStatus(store StatusStore) {
	profile := c.Profiles[c.Default.Name]
	statusMap := map[string]StatusStore{}

	if profile.StatusList != nil {
		statusMap = profile.StatusList
	}

	/*
		# Breakdown emoji and emoji text
			Default store.Emoji = 👋 :wave:
			splitted[0] = 👋
			splitted[1]  = :wave:

	*/
	splitted := strings.Fields(store.Emoji)
	store.Emoji = splitted[1]

	statusMapKey := fmt.Sprintf("%s  %s %s", splitted[0], color.Green.Sprintf(store.Status), color.Gray.Sprintf(store.UserDefinedDuration))
	statusMap[statusMapKey] = store

	c.Profiles[c.Default.Name] = Profile{Token: profile.Token, StatusList: statusMap}
	log.Printf("After appending %+v\n", c)
	c.Save()
}

func (c *Config) AddGlobalStatus(store StatusStore) {
	log.Println("globals", c.Globals)
	if len(c.Globals) == 0 {
		log.Println("Appending Globals")
		c.Globals = append(c.Globals, store)
		log.Println("Appended Globals", c.Globals)
		return
	}

	for k, v := range c.Globals {
		if store.Status == v.Status {
			replace := false
			prompt := &survey.Confirm{
				Message: "Global status exists: Update?",
			}
			survey.AskOne(prompt, &replace)
			if replace {
				c.Globals[k] = store
				return
			}
			return
		}
	}
	c.Globals = append(c.Globals, store)
}

func (c *Config) Save() {
	data, err := yaml.Marshal(c)
	utils.CheckIfError(err)
	err = ioutil.WriteFile(ConfigFilePath, data, 0777)
	utils.CheckIfError(err)
}

func (c *Config) DeleteProfile(profile string) {
	confirm := false
	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Are you sure you want to delete the profile: %s?", profile),
	}
	survey.AskOne(prompt, &confirm)
	fmt.Println("Profiles", c.Profiles)
	if confirm {
		delete(c.Profiles, profile)
		c.Save()
	}
}
