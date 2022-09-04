package handlers

import (
	"strings"

	"github.com/foxieze/tsundoku-server/config"
	"github.com/foxieze/tsundoku-server/entities"
	"github.com/foxieze/tsundoku-server/services"
	"github.com/gofiber/fiber/v2"
)

func GetBook(c *fiber.Ctx) error {
	var book entities.Book

	config.Database.Find(&book, c.Params("id"))

	// get token from header
	token := strings.Split(c.Get("Authorization"), " ")[1] // Bearer <token>

	err := services.AuthenticateAgainstId(token, book.Bookshelf.UserID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	return c.JSON(book)
}

func GetBooks(c *fiber.Ctx) error {
	var books []entities.Book

	bookshelfId := c.Params("id")

	bookshelf := entities.Bookshelf{}
	config.Database.Find(&bookshelf, bookshelfId)

	// get token from header
	token := strings.Split(c.Get("Authorization"), " ")[1] // Bearer <token>

	err := services.AuthenticateAgainstId(token, bookshelf.UserID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	config.Database.Where("bookshelf_id = ?", bookshelfId).Find(&books)

	return c.JSON(books)
}
