package entity

import (
	"gmauth/domain"

	"github.com/gofiber/fiber/v2"
)

type AuthEntity struct {
	authRepo domain.AuthRepository
}

func NewAuthEntity(r domain.AuthRepository) domain.AuthEntity {
	return &AuthEntity{authRepo: r}
}

func (e AuthEntity) GetLogin(ctx *fiber.Ctx) (domain.Auth, error) {
	return domain.Auth{}, nil
}

func (e AuthEntity) Authorize() {
	return
}
