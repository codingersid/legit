package routes

import (
	"github.com/gofiber/fiber/v2"
)

func RouterWeb(c *fiber.App) {
	// route web
	c.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
