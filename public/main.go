package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	legitConfig "github.com/codingersid/legit-cli/config"
	"github.com/codingersid/legit/config"
	"github.com/codingersid/legit/database/migrations"
	"github.com/codingersid/legit/database/seeders"
	"github.com/codingersid/legit/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// engine config
	env := legitConfig.LoadEnv()
	llog := legitConfig.InitLogger("logs")

	// tanpa database
	if env["APP_NO_DB"] == "false" {
		config.ConnectDB()
		// database config
		if env["APP_ENV"] == "local" {
			migrations.RunMigration()
			seeders.RunSeeder()
		}
	}

	// buat engine
	resourcesPath := "./resources"
	if _, err := os.Stat(resourcesPath); os.IsNotExist(err) {
		resourcesPath = "../resources"
	}
	engine := html.New(resourcesPath, ".html")

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

	// konfigurasi middleware
	app.Use(cors.New())

	// konfigurasi helmet
	csp := config.ConfigCSP
	app.Use(helmet.New(helmet.Config{
		ContentSecurityPolicy: csp(),
	}))

	// Middleware untuk memfilter double slash di akhir URL
	app.Use(func(c *fiber.Ctx) error {
		path := c.Path()
		if strings.HasSuffix(path, "//") {
			newPath := strings.TrimSuffix(path, "/")
			return c.Redirect(newPath)
		}
		return c.Next()
	})

	// Panggil routes web
	routes.RouterWeb(app)

	// lokasi file statics (images, css, js, plugins)
	staticPath := "./public/statics"
	if _, err := os.Stat(staticPath); os.IsNotExist(err) {
		staticPath = "statics"
	}
	app.Static("/", staticPath)

	// Kirim ke server
	errServe := app.Listen(env["APP_URL"] + ":" + env["APP_PORT"])
	if errServe != nil {
		fmt.Println("Error:", errServe)
	}

	// tutup log
	legitConfig.CloseLogger(llog)
}
