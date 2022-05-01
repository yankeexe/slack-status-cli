package utils

import (
	"github.com/gookit/color"
	"os"
)

func CheckIfError(err error) {
	if err == nil {
		return
	}
	color.Red.Println(err)
	os.Exit(1)
}

func HandleNoProfiles() {
	color.Red.Println(`No profiles found.

$ st profile --create
to create a new slack profile`)

	os.Exit(1)
}

// DurationConverter takes period and duration for a status and converts it into minutes
func DurationConverter(period int, duration string) int {
	switch duration {
	case "m", "min", "mins", "minutes":
		return period
	case "h", "hr", "hour", "hours":
		return period * 60
	case "d", "day", "days":
		return period * 1440
	}
	return 0
}
