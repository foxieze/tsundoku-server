package handlers

import (
	"github.com/foxieze/tsundoku-server/config"
	"github.com/foxieze/tsundoku-server/entities"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	var user entities.User

	config.Database.Find(&user, c.Params("id"))

	return c.JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	var users []entities.User

	config.Database.Find(&users)

	return c.JSON(users)
}

func CreateUser(c *fiber.Ctx) error {
	var user entities.User

	if err := c.BodyParser(&user); err != nil {
		return err
	}

	config.Database.Create(&user)

	return c.JSON(user)
}
