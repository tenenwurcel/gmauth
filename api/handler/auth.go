package handler

import (
	"encoding/hex"
	"fmt"
	"gmauth/config"
	"gmauth/domain"
	"gmauth/utils"
	"net/url"
	"time"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	authEntity domain.AuthEntity
}

type todoLater struct {
	Code  string
	State string
}

func NewAuthHandler(app *fiber.App) {
	handler := new(AuthHandler)

	app.Get("api/v1/auth/login", handler.GetLogin)
	app.Get("api/v1/auth/callback", handler.OAuthCallback)
	app.Post("api/v1/auth/authorize", handler.Authorize)
}

func (a *AuthHandler) GetLogin(c *fiber.Ctx) error {
	auth, err := a.authEntity.GetLogin(c)

	if err != nil {
		return err
	}

	c.Cookie(&fiber.Cookie{
		Name:     "sid",
		Value:    auth.SIDCrypt,
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   false,
		HTTPOnly: true,
		SameSite: "strict",
		Domain:   ".tenenwurcel.com.br",
	})

	discordOAuth := config.GetDiscordOAuth()
	authUrl := discordOAuth.GetDiscordAuthUrl(auth.State.String())

	return c.Send([]byte(authUrl))
}

func (a *AuthHandler) Authorize(c *fiber.Ctx) error {
	sid, err := hex.DecodeString(c.Cookies("sid"))
	if err != nil {
		fmt.Println(err)
	}

	cypher := utils.GetCypher()
	plain, err := cypher.Decrypt(sid)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(plain))
	return c.JSON("ok")
}

func (a *AuthHandler) OAuthCallback(c *fiber.Ctx) error {
	todoLater := new(todoLater)
	if err := c.QueryParser(todoLater); err != nil {
		return err
	}
	fmt.Println(todoLater.Code)

	params := &url.Values{}
	params.Add("code", todoLater.Code)
	params.Add("state", todoLater.State)

	return c.Redirect("https://tenenwurcel.com.br/")
}
