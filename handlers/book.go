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

	// get bookshelf from id
	bookshelf := entities.Bookshelf{}
	config.Database.Find(&bookshelf, book.BookshelfID)

	// get user from id
	user := entities.User{}
	config.Database.Find(&user, bookshelf.UserID)

	err := services.AuthenticateAgainstId(token, user.Id)
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

func CreateBook(c *fiber.Ctx) error {
	var book entities.Book

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

	if err := c.BodyParser(&book); err != nil {
		return err
	}

	book.BookshelfID = int(bookshelf.ID)

	config.Database.Create(&book)

	return c.JSON(book)
}

// get all books for user
func GetAllBooks(c *fiber.Ctx) error {
	var books []entities.Book

	// get token from header
	token := strings.Split(c.Get("Authorization"), " ")[1] // Bearer <token>

	userId, err := services.GetIdFromToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	bookshelves := []entities.Bookshelf{}
	config.Database.Where("user_id = ?", userId).Find(&bookshelves)

	for _, bookshelf := range bookshelves {
		var bookshelfBooks []entities.Book
		config.Database.Where("bookshelf_id = ?", bookshelf.ID).Find(&bookshelfBooks)
		books = append(books, bookshelfBooks...)
	}

	return c.JSON(books)
}
