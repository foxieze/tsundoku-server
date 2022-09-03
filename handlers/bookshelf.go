package handlers

import (
	"github.com/foxieze/tsundoku-server/config"
	"github.com/foxieze/tsundoku-server/entities"
	"github.com/gofiber/fiber/v2"
)

func GetBookshelf(c *fiber.Ctx) error {
	var bookshelf entities.Bookshelf

	config.Database.Find(&bookshelf, c.Params("id"))

	return c.JSON(bookshelf)
}

func GetBookshelves(c *fiber.Ctx) error {
	var bookshelves []entities.Bookshelf

	config.Database.Find(&bookshelves)

	return c.JSON(bookshelves)
}

func CreateBookshelf(c *fiber.Ctx) error {
	var bookshelf entities.Bookshelf

	if err := c.BodyParser(&bookshelf); err != nil {
		return err
	}

	config.Database.Create(&bookshelf)

	return c.JSON(bookshelf)
}
