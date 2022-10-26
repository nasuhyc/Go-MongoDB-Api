package app

import (
	"Go-MongoDb-Api/models"
	"Go-MongoDb-Api/services"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	Service services.UserService
}

func (h UserHandler) CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	result, err := h.Service.UserInsert(user)

	if err != nil || result.Status == false {
		return err
	}
	return c.Status(200).JSON(result)

}
