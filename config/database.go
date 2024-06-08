package config

import (
	"fmt"
	"log"
	"os"

	legitConfig "github.com/codingersid/legit-cli/config"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// db config
func loadDBConfig() legitConfig.DBConfig {
	env := legitConfig.LoadEnv()
	return legitConfig.DBConfig{
		DbUser:    env["DB_USER"],
		DbPass:    env["DB_PASS"],
		DbName:    env["DB_NAME"],
		DbHost:    env["DB_HOST"],
		DbPort:    env["DB_PORT"],
		DbCharset: env["DB_CHARSET"],
		DbDriver:  env["DB_DRIVER"],
	}
}

// general init db
func ConnectDB() {
	dbConfig := loadDBConfig()
	legitConfig.ConnectDB(dbConfig)
}

// custom init db
func InitDB() *gorm.DB {
	var err error
	var DSN string
	dbConfig := loadDBConfig()
	switch dbConfig.DbDriver {
	case "mysql":
		DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			dbConfig.DbUser, dbConfig.DbPass, dbConfig.DbHost, dbConfig.DbPort, dbConfig.DbName, dbConfig.DbCharset)
		DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	case "postgres":
		DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbConfig.DbHost, dbConfig.DbUser, dbConfig.DbPass, dbConfig.DbName, dbConfig.DbPort)
		DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	default:
		return nil
	}
	if err != nil {
		fmt.Printf("Database connection error: %v\n", err)
		return nil
	}
	return DB
}

// custom init db dengan error tercatat ke log, tidak di print di terminal
func InitDBWithoutError() *gorm.DB {
	var err error
	var DSN string
	dbConfig := loadDBConfig()
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel: logger.Silent, // Menonaktifkan semua log
		},
	)

	switch dbConfig.DbDriver {
	case "mysql":
		DSN = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			dbConfig.DbUser, dbConfig.DbPass, dbConfig.DbHost, dbConfig.DbPort, dbConfig.DbName, dbConfig.DbCharset)
		DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{
			Logger: newLogger,
		})
	case "postgres":
		DSN = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			dbConfig.DbHost, dbConfig.DbUser, dbConfig.DbPass, dbConfig.DbName, dbConfig.DbPort)
		DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{
			Logger: newLogger,
		})
	default:
		return nil
	}
	if err != nil {
		log.Printf("Database connection error: %v", err)
		return nil
	}
	return DB
}
