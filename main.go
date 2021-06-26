package main

import (
	"fmt"
	_handler "gmauth/api/handler"
	"gmauth/config"
	"gmauth/utils"
)

func main() {
	utils.SetupCypher()
	config.SetupDiscordOauth()
	config.SetupApp()
	config.RegisterFiberHandler(_handler.NewAuthHandler)

	app := config.GetApp()
	err := app.Listen(config.GetAppPort())
	if err != nil {
		fmt.Println(err)
	}
}
