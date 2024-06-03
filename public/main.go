package main

import (
	"fmt"
	"os"
	"strconv"

	legitConfig "github.com/codingersid/legit-cli/config"
	"github.com/codingersid/legit/config"
	"github.com/codingersid/legit/database/migrations"
	"github.com/codingersid/legit/database/seeders"
	"github.com/codingersid/legit/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	/*
		|	Konfigurasi ENV
		|	Untuk membaca file .env
	*/
	env := legitConfig.LoadEnv()

	/*
		|	Konfigurasi logrus
		|	Memanggil fungsi log
	*/
	llog := legitConfig.InitLogger("logs")

	/*
		|	Konfigurasi Database
		|	Jika APP_NO_DB pada .env false, maka mengkoneksikan database dan wajib ada database
	*/
	if env["APP_NO_DB"] == "false" {
		config.ConnectDB()
		/*
			|	Jika APP_ENV pada .env local, maka setiap server dijalankan, database akan menajalankan migration dan seeder
		*/
		if env["APP_ENV"] == "local" {
			migrations.RunMigration()
			seeders.RunSeeder()
		}
	}

	/*
		|	Konfigurasi Engine Template
	*/
	resourcesPath := "./resources"
	if _, err := os.Stat(resourcesPath); os.IsNotExist(err) {
		resourcesPath = "../resources"
	}
	engine := html.New(resourcesPath, ".html")

	/*
		|	Konfigurasi Menjalankan Fiber
	*/
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

	/*
		|	Konfigurasi Middleware
	*/
	config.RunConfigApps(app)

	/*
		|	Konfigurasi Route
		|	API Router
	*/
	// REST_API
	if env["REST_API"] == "true" {
		routes.RouterApi(app)
	}

	/*
		|	Konfigurasi Route
		|	Web Router
	*/
	routes.RouterWeb(app)

	/*
		|	Konfigurasi File Static
		|	lokasi file statics (images, css, js, plugins)
	*/
	staticPath := "./public/statics"
	if _, err := os.Stat(staticPath); os.IsNotExist(err) {
		staticPath = "statics"
	}
	app.Static("/", staticPath)
	// statics dengan FileSystem
	// config.ConfigFileSystem(app)

	/*
		|	Menajalankan Server
		|	APP_URL pada .env untuk menentukan URL web
		|	APP_PORT pada .env untuk menentukan PORT yang digunakan pada web
	*/
	errServe := app.Listen(env["APP_URL"] + ":" + env["APP_PORT"])
	if errServe != nil {
		fmt.Println("Error:", errServe)
	}

	/*
		|	Konfigurasi logrus
		|	Menutup fungsi log
	*/
	legitConfig.CloseLogger(llog)
}
