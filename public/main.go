package main

import (
	"fmt"
	"strconv"

	legitConfig "github.com/codingersid/legit-cli/config"
	"github.com/codingersid/legit/config"
	"github.com/codingersid/legit/database/migrations"
	"github.com/codingersid/legit/database/seeders"
	"github.com/codingersid/legit/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// engine config
	env := legitConfig.LoadEnv()
	llog := legitConfig.InitLogger("logs")
	config.ConnectDB()

	// database config
	if env["APP_ENV"] == "local" {
		migrations.RunMigration()
		seeders.RunSeeder()
	}

	// buat engine
	engine := html.New("./resources", ".html")

	// panggil Fiber
	app := fiber.New(fiber.Config{
		Views:   engine,
		AppName: env["APP_NAME"],
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			message := "Internal Server Error"
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				message = e.Message
			}

			llog.WithError(err).WithField("code", code).WithField("message", message).Error("Internal Server Error")
			return c.Render("errors/errors", fiber.Map{
				"Title":   strconv.Itoa(code) + " - " + message,
				"Code":    code,
				"Message": message,
			})
		},
	})

	// lokasi file statics (images, css, js, plugins)
	app.Static("/", "./public/statics")

	// konfigurasi middleware
	app.Use(cors.New())

	// Panggil routes web
	routes.RouterWeb(app)

	// Kirim ke server
	errServe := app.Listen(env["APP_URL"] + ":" + env["APP_PORT"])
	if errServe != nil {
		fmt.Println("Error:", errServe)
	}

	// tutup log
	legitConfig.CloseLogger(llog)
}
