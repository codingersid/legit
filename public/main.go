package main

import (
	"fmt"
	"os"
	"path/filepath"

	legitConfig "github.com/codingersid/legit-cli/config"
	"github.com/codingersid/legit/config"
	"github.com/codingersid/legit/database/migrations"
	"github.com/codingersid/legit/database/seeders"
	"github.com/codingersid/legit/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"
)

func main() {
	env := legitConfig.LoadEnv()

	// Buat file log
	file := CreateLogFile()
	defer file.Close()

	// Set file sebagai output log
	logrus.SetOutput(file)

	// panggil database
	config.ConnectDB()

	if env["APP_ENV"] == "local" {
		// panggil migration
		migrations.RunMigration()

		// panggil seeder
		seeders.RunSeeder()
	}

	// panggil Fiber
	app := fiber.New(fiber.Config{
		AppName: env["APP_NAME"],
	})

	// Middleware
	app.Use(cors.New())

	// Panggil routes web
	routes.RouterWeb(app)

	// Exceptions

	// Kirim ke server
	errServe := app.Listen(env["APP_URL"] + ":" + env["APP_PORT"])
	if errServe != nil {
		fmt.Println("Error:", errServe)
	}
}

func CreateLogFile() *os.File {
	// Path ke file log, relative terhadap working directory aplikasi
	logFilePath := filepath.Join("..", "logs", "logs.log")

	// Buat file log di lokasi yang ditentukan
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		// Jika file log tidak dapat dibuat di lokasi yang diinginkan, gunakan lokasi dari root aplikasi
		rootPath, err := os.Getwd()
		if err != nil {
			logrus.Fatal(err)
		}
		logFilePath = filepath.Join(rootPath, "logs", "logs.log")
		file, err = os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			logrus.Fatal(err)
		}
	}

	return file
}
