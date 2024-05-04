package routes

import (
	legitConfig "github.com/codingersid/legit-cli/config"
	"github.com/gofiber/fiber/v2"
)

func RouterWeb(c *fiber.App) {
	// route web
	c.Get("/", legitConfig.Views("welcome", fiber.Map{}))
}
