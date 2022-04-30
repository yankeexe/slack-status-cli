package utils

import (
	"fmt"
	"os"
	"regexp"

	"github.com/gookit/color"
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
func ParseDuration(time string) {
	r := regexp.MustCompile(`(?P<Period>^\d+).?(?P<Duration>m$|min$|minutes?|h$|hr$|hours?|d$|days?)?`)
	match := r.FindStringSubmatch(time)

	if len(match) == 0 {
		color.Red.Println(`Invalid status duration
Use 'st help add' for more information.
		`)
		os.Exit(1)
	}

	statusPeriod := match[1]
	if match[2] == "" {
		match[2] = "minutes"
	}
	fmt.Println("status period:", statusPeriod)
	fmt.Println("status duration", match[2])

}
