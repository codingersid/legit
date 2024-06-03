package config

import (
	"github.com/gofiber/fiber/v2/middleware/redirect"
)

// Konfigurasi untuk middleware REDIRECT
var ConfigRedirect = redirect.Config{
	Rules:      getRedirect(false),
	StatusCode: 301,
}

func getRedirect(status bool) map[string]string {
	if !status {
		return map[string]string{}
	}
	return map[string]string{
		"/contoh":        "/controller",
		"/contoh/edit/*": "/controller/edit/$1",
	}
}
