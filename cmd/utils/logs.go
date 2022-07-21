package utils

import (
	"github.com/gookit/color"
	"os"
)

func SendLog(typeLog string, message string) {
	switch typeLog {
	case "error":
		color.Println(`<red>[ERROR]</> `, message)
		os.Exit(1)

	case "server":
		color.Println(`<yellow>[SERVER]</> `, message)

	case "system":
		color.Println(`<yellow>[SYSTEM]</> `, message)

	}
}
