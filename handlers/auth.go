package handlers

// Import fiber and services
import (
	"github.com/foxieze/tsundoku-server/services"
	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	// Get username and password from request body
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Get token from auth service
	token, err := services.LoginUser(username, password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// Return token + info
	return c.JSON(fiber.Map{
		"token":    token,
		"username": c.FormValue("username"),
	})
}

func Register(c *fiber.Ctx) error {
	// Get username and password from request body
	username := c.FormValue("username")
	password := c.FormValue("password")

	// Get token from auth service
	token, err := services.RegisterUser(username, password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// Return token + info
	return c.JSON(fiber.Map{
		"token":    token,
		"username": c.FormValue("username"),
	})
}
