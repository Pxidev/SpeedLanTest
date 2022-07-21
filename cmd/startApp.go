package cmd

import (
	"core/cmd/vars"
	"core/src/core"
	"github.com/gookit/color"
)

func StartApp() {

	color.Println(`<red>      
 __                     _   __            _____          _   
/ _\_ __   ___  ___  __| | / /  __ _ _ __/__   \___  ___| |_ 
\ \| '_ \ / _ \/ _ \/ _´ |/ /  / _´ | '_ \ / /\/ _ \/ __| __|
_\ \ |_) |  __/  __/ (_| / /__| (_| | | | / / |  __/\__ \ |_ 
\__/ .__/ \___|\___|\__,_\____/\__,_|_| |_\/   \___||___/\__|
   |_|</> 
--------------------------------------------------------------
		<cyan>Created by Pixidev</> - <yellow>v` + vars.Version + `</>
	`)

	Port := getPort()

	server := core.Server{
		IP:      "0.0.0.0",
		Port:    Port,
		Version: vars.Version,
	}

	server.Run()

}
