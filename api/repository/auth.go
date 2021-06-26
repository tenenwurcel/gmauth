package repository

import (
	"gmauth/domain"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type AuthRepository struct {
	client *http.Client
}

func NewAuthEntity(r domain.AuthRepository) domain.AuthEntity {
	return &AuthRepository{client: nil}
}

func (r AuthRepository) GetLogin(ctx *fiber.Ctx) (domain.Auth, error) {
	return domain.Auth{}, nil
}

func (r AuthRepository) Authorize() {
	return
}
