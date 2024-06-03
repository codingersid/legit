package config

import "github.com/gofiber/fiber/v2/middleware/rewrite"

// Konfigurasi untuk middleware REWRITE
var ConfigRewrite = rewrite.Config{
	Rules: getRewrite(false),
}

func getRewrite(status bool) map[string]string {
	if !status {
		return map[string]string{}
	}
	return map[string]string{
		"/contoh":        "/controller",
		"/contoh/edit/*": "/controller/edit/$1",
	}
}
