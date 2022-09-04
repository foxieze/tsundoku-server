package main

import (
	"github.com/foxieze/tsundoku-server/config"
	"github.com/foxieze/tsundoku-server/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.Connect()

	// Bookshelves
	app.Get("/bookshelves", handlers.GetBookshelves)
	app.Get("/bookshelves/:id", handlers.GetBookshelf)
	app.Post("/bookshelves", handlers.CreateBookshelf)

	// Books
	app.Get("/books/:id", handlers.GetBooks)
	app.Get("/book/:id", handlers.GetBook)
	app.Post("/books/:id", handlers.CreateBook)

	// Users
	app.Post("/login", handlers.Login)
	app.Post("/register", handlers.Register)

	app.Listen(":3000")
}
