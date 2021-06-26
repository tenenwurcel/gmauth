package config

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var app *fiber.App
var port string

type handle func(*fiber.App)

func SetupApp() {
	App := fiber.New()
	fmt.Println("ok")
	App.Use(cors.New(cors.Config{
		AllowOrigins:     "https://tenenwurcel.com.br",
		AllowHeaders:     "*",
		AllowCredentials: true,
	}))
	setApp(App)
	setPort("80")
}

func RegisterFiberHandler(hdl handle) {
	hdl(app)
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
