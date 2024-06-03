package config

import (
	"net/http"
	"os"
	"strings"
	"time"

	legitConfig "github.com/codingersid/legit-cli/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/expvar"
	"github.com/gofiber/fiber/v2/middleware/favicon"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/idempotency"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/redirect"
	"github.com/gofiber/fiber/v2/middleware/rewrite"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Mulai Session
var StartSession = session.New(ConfigSession)

/*
|	Konfigurasi App
|	Segala konfigurasi framework untuk app.Use
*/
func RunConfigApps(app *fiber.App) {
	env := legitConfig.LoadEnv()
	pathSlashRemoval(app)
	// recover untuk panic handler
	app.Use(recover.New())
	// ENCRYPT COOKIE
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: encryptcookie.GenerateKey(),
	}))
	// idempotency
	app.Use(idempotency.New())
	// Limitter | Melindungi dari DDOS Attack
	configLimiter(app)
	// ETag
	configETag(app)
	// Favicon
	configFavicon(app)
	// csrf
	configCsrf(app)
	// configRedirect
	configRedirect(app)
	// configRewrite
	configRewrite(app)
	// healthcheck
	app.Use(healthcheck.New(healthcheck.Config{
		LivenessEndpoint:  "/cek-live",
		ReadinessEndpoint: "/cek-ready",
	}))
	// ExpVar
	if env["APP_ENV"] == "local" {
		// cek debug hanya saat Development
		app.Use("/debug/vars", expvar.New())
	}
	// SEC_CACHE
	if env["SEC_CACHE"] == "true" {
		configCache(app)
	}
	// SEC_COMPRESS
	if env["SEC_COMPRESS"] == "true" {
		configCompress(app)
	}
	// METRICS | CEK PERFORMA APLIKASI
	app.Get("/cek-performa", monitor.New(monitor.Config{Title: "Data Performa Aplikasi : " + env["APP_NAME"]}))
	// helmet
	configHelmet(app)
	// CORS
	configCors(app)
}

/*
| Konfigurasi untuk menghapus slash (/) diakhir URI
*/
func pathSlashRemoval(app *fiber.App) {
	app.Use(func(c *fiber.Ctx) error {
		url := c.Path()
		if url != "/" && strings.HasSuffix(url, "/") {
			url = strings.TrimSuffix(url, "/")
			return c.Redirect(url)
		}
		return c.Next()
	})
}

/*
| Konfigurasi untuk ETag
*/
func configLimiter(app *fiber.App) {
	// Middleware untuk ETag, akan di-skip jika route yang diakses adalah /api
	app.Use(func(c *fiber.Ctx) error {
		if strings.HasPrefix(c.Path(), "/api") {
			return limiter.New(limiter.Config{
				Max:        20,               // Jumlah maksimum permintaan per durasi
				Expiration: 30 * time.Second, // Durasi dalam detik sebelum reset
				KeyGenerator: func(c *fiber.Ctx) string {
					return c.IP() // Menggunakan IP klien sebagai kunci
				},
				LimitReached: func(c *fiber.Ctx) error {
					return c.Status(fiber.StatusTooManyRequests).SendString("Too many requests, please try again later.")
				},
			})(c)
		}
		return limiter.New(limiter.Config{
			Max:        20,               // Jumlah maksimum permintaan per durasi
			Expiration: 30 * time.Second, // Durasi dalam detik sebelum reset
			KeyGenerator: func(c *fiber.Ctx) string {
				return c.IP() // Menggunakan IP klien sebagai kunci
			},
			LimitReached: func(c *fiber.Ctx) error {
				return c.Render("errors/errors", fiber.Map{
					"Title":   "429 - Too many requests, please try again later.",
					"Code":    "429",
					"Message": "Too many requests, please try again later.",
				})
			},
		})(c)
	})
}

/*
| Konfigurasi untuk ETag
*/
func configETag(app *fiber.App) {
	// Middleware untuk ETag, akan di-skip jika route yang diakses adalah /api
	app.Use(func(c *fiber.Ctx) error {
		if strings.HasPrefix(c.Path(), "/api") {
			return c.Next()
		}
		return etag.New()(c)
	})
}

/*
| Konfigurasi untuk favicon
*/
func configFavicon(app *fiber.App) {
	// Middleware untuk favicon, akan di-skip jika route yang diakses adalah /api
	app.Use(func(c *fiber.Ctx) error {
		if strings.HasPrefix(c.Path(), "/api") {
			return c.Next()
		}
		staticPath := "./public/statics"
		if _, err := os.Stat(staticPath); os.IsNotExist(err) {
			staticPath = "statics"
		}
		return favicon.New(favicon.Config{
			File: staticPath + "/favicon.ico",
			URL:  "/favicon.ico",
		})(c)
	})
}

/*
| Konfigurasi untuk FileSystem
*/
func ConfigFileSystem(app *fiber.App) {
	// Middleware untuk FileSystem, akan di-skip jika route yang diakses adalah /api
	app.Use(func(c *fiber.Ctx) error {
		if strings.HasPrefix(c.Path(), "/api") {
			return c.Next()
		}
		staticPath := "./public/statics"
		if _, err := os.Stat(staticPath); os.IsNotExist(err) {
			staticPath = "statics"
		}
		return filesystem.New(filesystem.Config{
			Root:         http.Dir(staticPath),
			Browse:       false,
			Index:        "index.html",
			NotFoundFile: "404.html",
			MaxAge:       3600,
		})(c)
	})
}

/*
| Konfigurasi untuk Helmet
| Ini membantu digunakan untuk meningkatkan keamanan aplikasi web dengan mengatur berbagai header HTTP yang melindungi dari serangan umum di web.
*/
func configHelmet(app *fiber.App) {
	app.Use(helmet.New(ConfigHelmet))
}

/*
| Konfigurasi untuk cache
| Matikan konfigurasi ini di .env jika website yang membutuhkan data realtime
*/
func configCache(app *fiber.App) {
	// Middleware untuk cache, akan di-skip jika route yang diakses adalah /api
	app.Use(func(c *fiber.Ctx) error {
		if strings.HasPrefix(c.Path(), "/api") {
			return c.Next()
		}
		return cache.New(ConfigCache)(c)
	})
}

/*
| Konfigurasi untuk Compress
| Ini membantu mengurangi ukuran respons yang dikirimkan dari server ke klien, yang dapat menghemat bandwidth dan mempercepat waktu pemuatan halaman.
| Level: LevelDisabled, LevelDefault, LevelBestSpeed, LevelBestCompression
*/
func configCompress(app *fiber.App) {
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))
}

/*
| Konfigurasi untuk CORS
| Untuk mengatur header CORS pada respons HTTP yang dihasilkan oleh server.
| Ini memungkinkan atau membatasi akses dari sumber daya yang berada di domain yang berbeda.
*/
func configCors(app *fiber.App) {
	app.Use(cors.New(ConfigCors))
}

/*
| Konfigurasi untuk CSRF
| Untuk melindungi aplikasi web dari serangan CSRF (Cross-Site Request Forgery) dengan menambahkan token CSRF ke setiap formulir atau permintaan yang memerlukan otorisasi.
*/
func configCsrf(app *fiber.App) {
	// Middleware untuk CSRF, akan di-skip jika route yang diakses adalah /api
	app.Use(func(c *fiber.Ctx) error {
		if strings.HasPrefix(c.Path(), "/api") {
			return c.Next()
		}
		return csrf.New(ConfigCsrf)(c)
	})
}

/*
| Konfigurasi untuk REDIRECT
| Untuk redirect url tertentu secara permanent
*/
func configRedirect(app *fiber.App) {
	// Middleware untuk redirect, akan di-skip jika route yang diakses adalah /api
	app.Use(func(c *fiber.Ctx) error {
		if strings.HasPrefix(c.Path(), "/api") {
			return c.Next()
		}
		return redirect.New(ConfigRedirect)(c)
	})
}

/*
| Konfigurasi untuk REWRITE
| Untuk rewrite url tertentu ke url lain yang ada di route
*/
func configRewrite(app *fiber.App) {
	// Middleware untuk rewrite, akan di-skip jika route yang diakses adalah /api
	app.Use(func(c *fiber.Ctx) error {
		return rewrite.New(ConfigRewrite)(c)
	})
}
