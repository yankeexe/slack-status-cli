package config

type Config struct {
	Globals  []string
	Default  string
	Profiles map[string]Profile
}

type Profile struct {
	Name   string   `mapstructure:"name"`
	Token  string   `mapstructure:"token"`
	Status []string `mapstructure:"Status"`
}

func (c *Config) GetProfiles() []string {
	profiles := []string{}
	for k, _ := range c.Profiles {
		profiles = append(profiles, k)
	}
	return profiles
}
