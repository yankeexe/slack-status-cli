package utils

import (
	"os"

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
