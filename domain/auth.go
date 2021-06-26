package domain

import (
	"encoding/hex"
	"gmauth/utils"

	"github.com/gofiber/fiber/v2"
	uuid "github.com/satori/go.uuid"
)

type Auth struct {
	SID      uuid.UUID `json:"-"`
	State    uuid.UUID `json:"state"`
	SIDCrypt string    `json:"sid"`
}

func NewAuth() (Auth, error) {
	sid := uuid.NewV4()
	state := uuid.NewV4()

	cypher := utils.GetCypher()
	ciphertext, err := cypher.Encrypt([]byte(sid.String()))
	if err != nil {
		return Auth{}, err
	}

	return Auth{SIDCrypt: hex.EncodeToString(ciphertext), State: state, SID: sid}, nil
}

type AuthEntity interface {
	GetLogin(ctx *fiber.Ctx) (Auth, error)
	Authorize()
	//getDiscordUser()
	//checkGuildAuth()
}

type AuthRepository interface {
	GetLogin(ctx *fiber.Ctx) (Auth, error)
	Authorize()
}
