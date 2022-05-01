package config

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"io/ioutil"
	"log"
	"math"
	"os"
	"os/user"
	"regexp"
	"strconv"

	"path/filepath"

	"github.com/spf13/viper"

	"github.com/gookit/color"
	"github.com/yankeexe/slack-status-cli/internal/utils"
)

var usr, _ = user.Current()

var ConfigDirPath = filepath.Join(usr.HomeDir, ".slack-status-cli")
var ConfigFilePath = filepath.Join(ConfigDirPath, "config")

type Config struct {
	Default  ProfileInfo
	Globals  []string
	Profiles map[string]Profile
}

type ProfileInfo struct {
	Name  string `mapstructure:"name"`
	Token string `mapstructure:"token"`
}

type Profile struct {
	Token      string `mapstructure:"token"`
	StatusList []StatusStore
}

type StatusStore struct {
	Status              string
	UserDefinedDuration string
	Period              int
	Emoji               string
}

/*
	Parses user duration for minute, hour, and day.
	Defaults to minutes

	Check the duration string has one of the suffix  (minutes, hour, days) in a array.
	If it has value in that array then take that suffix and get numbers our of it.

	@CONDITIONS
	minutes cannot be more than 43800
	hours cannot be more than 730
	day cannot be more than 30

	@REGEX:
		(?P<Name>pattern) = match group
		^\d+ = starts with any digits with one or more matches
		.? = optional any character after the digits
		---
		-- Matches for duration suffix
		(m$|min$|minutes?|h$|hr$|hours?|d$|days?)?
		$ = match with character/word ending before the symbol
		| = OR
*/
func (status *StatusStore) ParseDuration() {
	r := regexp.MustCompile(`(?P<Period>^\d+).?(?P<Duration>m$|min$|minutes?|h$|hr$|hours?|d$|days?)?`)
	match := r.FindStringSubmatch(status.UserDefinedDuration)
	log.Println("Match", match)

	if len(match) == 0 {
		color.Red.Println(`Invalid status duration
Use 'st help add' for more information.
		`)
		os.Exit(1)
	}

	statusPeriod, err := strconv.ParseFloat(match[1], 32)
	utils.CheckIfError(err)

	// statusPeriod will be floored with ``int`` type casting.
	normalizedStatusPeriod := int(math.Abs(statusPeriod))

	statusDuration := match[2]
	if statusDuration == "" {
		statusDuration = "minutes"
	}
	absPeriod := utils.DurationConverter(normalizedStatusPeriod, statusDuration)
	status.UserDefinedDuration = fmt.Sprintf("%d %s", normalizedStatusPeriod, statusDuration)
	status.Period = absPeriod
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
	c.Profiles[c.Default.Name] = Profile{Token: profile.Token, StatusList: append(profile.StatusList, store)}
	log.Printf("After appending %+v\n", c)
	c.Save()
}

func (c *Config) Save() {
	data, err := toml.Marshal(c)
	utils.CheckIfError(err)
	err = ioutil.WriteFile(ConfigFilePath, data, 0777)
	utils.CheckIfError(err)
}
