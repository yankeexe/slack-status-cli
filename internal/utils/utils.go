package utils

import (
	"fmt"
	"github.com/gookit/color"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
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
// also returns absolute duration for showing on help messages
func DurationConverter(period int, duration string) (int, string) {
	var du string
	switch duration {
	case "m", "min", "mins", "minutes":
		if du = "minute"; period > 1 {
			du = "minutes"
		}
		return period, du
	case "h", "hr", "hour", "hours":
		if du = "hour"; period > 1 {
			du = "hours"
		}
		return period * 60, du
	case "d", "day", "days":
		if du = "day"; period > 1 {
			du = "days"
		}
		return period * 1440, du
	}
	return 0, ""
}

type ParsedDuration struct {
	UserDefinedDuration string
	AbsolutePeriod      int
}

/*
	Parses user duration for minute, hour, and day.
	Defaults to minutes.

	@CONDITIONS
	minutes cannot be more than 43800
	hours cannot be more than 730
	day cannot be more than 30

	@REGEX:
		(?P<Name>pattern) = match group
		^\d+ = starts with any digits with one or more matches
		\s? = optional whitespace characters
		---
		-- Matches for duration suffix
		(m$|min$|minutes?|h$|hr$|hours?|d$|days?)?
		$ = match with character/word ending before the symbol
		| = OR
*/
func ParseDuration(duration string) ParsedDuration {
	fmt.Println("Got", duration)
	r := regexp.MustCompile(`(?P<Period>^\d+)\s?(?P<Duration>m$|min$|minutes?|h$|hr$|hours?|d$|days?)?`)
	match := r.FindStringSubmatch(duration)
	log.Println("Match", match)

	if len(match) == 0 {
		color.Red.Println(`Invalid status duration
Use 'st help add' for more information.
		`)
		os.Exit(1)
	}

	statusPeriod, err := strconv.ParseFloat(match[1], 32)
	CheckIfError(err)

	// statusPeriod will be floored with ``int`` type casting.
	normalizedStatusPeriod := int(math.Abs(statusPeriod))

	statusDuration := match[2]
	if statusDuration == "" {
		statusDuration = "minutes"
	}
	absPeriod, absDuration := DurationConverter(normalizedStatusPeriod, statusDuration)
	userDefinedDuration := fmt.Sprintf("%d %s", normalizedStatusPeriod, absDuration)
	return ParsedDuration{UserDefinedDuration: userDefinedDuration, AbsolutePeriod: absPeriod}
}
