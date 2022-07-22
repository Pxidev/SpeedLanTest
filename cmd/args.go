package cmd

import (
	"SpeedLanTest/cmd/utils"
	"os"
)

func getPort() string {
	args := os.Args[1]

	if !utils.IsNumber(args) {
		utils.SendLog("error", "Es necesario introducir un puerto formato numerico")
		return ""
	}

	return args
}
