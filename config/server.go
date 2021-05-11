package config

import (
	"fmt"
	_handlers "gmauth/api/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var app *fiber.App
var port string

func SetupApp() {
	App := fiber.New()
	App.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))
	registerHandlers(App)
	setApp(App)
	setPort("80")
}

func registerHandlers(r *fiber.App) {
	_handlers.NewAuthHandler(r)
}

func setApp(App *fiber.App) {
	app = App
}

func setPort(Port string) {
	port = Port
}

func GetApp() *fiber.App {
	return app
}

func GetAppPort() string {
	if port == "" {
		return ":8080"
	}
	return fmt.Sprintf(":%s", port)
}
