package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RouterApi(c *fiber.App) {
	// route api
	api := c.Group("/api")
	api.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  fiber.StatusOK,
			"message": "Selamat Datang di REST API LEGIT FRAMEWORK",
		})
	})
}
