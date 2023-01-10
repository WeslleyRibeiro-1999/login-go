package api

import (
	"github.com/WeslleyRibeiro-1999/login-go/models"
	"github.com/WeslleyRibeiro-1999/login-go/src/repository"
	"github.com/gofiber/fiber/v2"
)

type Login interface {
	CreateUser(c *fiber.Ctx) error
}

type login struct {
	repository repository.Repository
}

var _ Login = (*login)(nil)

func NewHandler(repo repository.Repository) Login {
	return &login{
		repository: repo,
	}
}

func (l *login) CreateUser(c *fiber.Ctx) error {
	var req *models.User

	if err := c.BodyParser(req); err != nil {
		return c.JSON(fiber.Map{"error": "erro ao tentar processar dados inseridos"})
	}

	// newUser, err := h.repository.SingUp(req)

	return nil
}
