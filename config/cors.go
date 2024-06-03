package config

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// Konfigurasi untuk middleware CORS
var ConfigCors = cors.Config{
	// Middleware selanjutnya dalam rantai middleware, nil berarti tidak ada middleware selanjutnya setelah CORS.
	Next: nil,
	// Fungsi untuk mengizinkan origin-origin tertentu. Jika nil, maka AllowOrigins akan digunakan.
	AllowOriginsFunc: nil,
	// Daftar origin yang diizinkan. Jika AllowOriginsFunc nil, ini akan digunakan.
	AllowOrigins: allowOrigins(true),
	// Metode HTTP yang diizinkan untuk permintaan CORS.
	AllowMethods: strings.Join([]string{
		fiber.MethodGet,
		fiber.MethodPost,
		fiber.MethodHead,
		fiber.MethodPut,
		fiber.MethodDelete,
		fiber.MethodPatch,
	}, ","),
	// Header HTTP yang diizinkan untuk permintaan CORS.
	AllowHeaders: allowHeaders(true),
	// Mengizinkan kredensial untuk dikirim atau tidak dalam permintaan CORS.
	AllowCredentials: true,
	// Header HTTP yang akan diungkapkan kepada klien.
	ExposeHeaders: "Content-Length, Content-Type",
	// Durasi dalam detik sebelum cache preflight CORS dianggap kadaluarsa.
	MaxAge: 86400,
}

// Fungsi untuk menghasilkan string yang berisi daftar origin yang diizinkan.
func allowOrigins(status bool) string {
	if !status {
		return "*"
	}
	slice := []string{
		"https://codingers.id",
		"https://source.unsplash.com",
	}
	return strings.Join(slice, ",")
}

// Fungsi untuk menghasilkan string yang berisi daftar header yang diizinkan.
func allowHeaders(status bool) string {
	if !status {
		return ""
	}
	slice := []string{
		"Origin",
		"Content-Type",
		"Accept",
		"Authorization",
	}
	return strings.Join(slice, ",")
}
