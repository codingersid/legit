package config

import (
	legitConfig "github.com/codingersid/legit-cli/config"
)

func ConnectDB() {
	env := legitConfig.LoadEnv()
	dbConfig := legitConfig.DBConfig{
		DbUser:    env["DB_USER"],
		DbPass:    env["DB_PASS"],
		DbName:    env["DB_NAME"],
		DbHost:    env["DB_HOST"],
		DbPort:    env["DB_PORT"],
		DbCharset: env["DB_CHARSET"],
		DbDriver:  env["DB_DRIVER"],
	}
	legitConfig.ConnectDB(dbConfig)
}
