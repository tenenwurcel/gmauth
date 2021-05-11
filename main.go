package main

import (
	"fmt"
	"gmauth/config"
	"gmauth/utils"
)

func main() {
	utils.SetupCypher()
	config.SetupDiscordOauth()
	config.SetupApp()
	app := config.GetApp()
	err := app.Listen(config.GetAppPort())
	if err != nil {
		fmt.Println(err)
	}
}
