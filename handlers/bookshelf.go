package handlers

import (
	"strings"

	"github.com/foxieze/tsundoku-server/config"
	"github.com/foxieze/tsundoku-server/entities"
	"github.com/foxieze/tsundoku-server/services"
	"github.com/gofiber/fiber/v2"
)

func GetBookshelf(c *fiber.Ctx) error {
	var bookshelf entities.Bookshelf

	config.Database.Find(&bookshelf, c.Params("id"))

	// get token from header
	token := strings.Split(c.Get("Authorization"), " ")[1] // Bearer <token>

	err := services.AuthenticateAgainstId(token, bookshelf.UserID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	return c.JSON(bookshelf)
}

func GetBookshelves(c *fiber.Ctx) error {
	var bookshelves []entities.Bookshelf

	userId, err := services.GetIdFromToken(strings.Split(c.Get("Authorization"), " ")[1])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	config.Database.Where("user_id = ?", userId).Find(&bookshelves)

	return c.JSON(bookshelves)
}

func CreateBookshelf(c *fiber.Ctx) error {
	var bookshelf entities.Bookshelf

	userId, err := services.GetIdFromToken(strings.Split(c.Get("Authorization"), " ")[1])
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	if err := c.BodyParser(&bookshelf); err != nil {
		return err
	}

	bookshelf.UserID = userId

	config.Database.Create(&bookshelf)

	return c.JSON(bookshelf)
}
