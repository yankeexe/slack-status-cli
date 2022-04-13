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
