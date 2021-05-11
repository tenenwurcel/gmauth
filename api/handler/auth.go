package handler

import (
	"fmt"
	"gmauth/domain"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct{}

func NewAuthHandler(app *fiber.App) {
	handler := &AuthHandler{}

	app.Get("api/v1/auth/login", handler.GetLogin)
	app.Get("/test", handler.Test)
}

func (a *AuthHandler) Test(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "randomvalue",
		Expires:  time.Now().Add(24 * time.Hour),
		SameSite: "none",
		Secure:   true,
	})
	return c.JSON("ok")
}

func (a *AuthHandler) GetLogin(c *fiber.Ctx) error {
	res, err := domain.NewAuth()
	if err != nil {
		// @TODO - return error
		return err
	}

	fmt.Printf("origin=>%s", c.Get("Origin"))

	c.Cookie(&fiber.Cookie{
		Name:     "token",
		Value:    "randomvalue",
		Expires:  time.Now().Add(24 * time.Hour),
		HTTPOnly: true,
		SameSite: "none",
		Secure:   true,
	})

	return c.JSON(res)
}
