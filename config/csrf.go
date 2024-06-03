package config

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"time"

	legitConfig "github.com/codingersid/legit-cli/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// Konfigurasi untuk middleware CSRF
var ConfigCsrf = csrf.Config{
	// KeyLookup menentukan lokasi di mana token CSRF akan dicari dalam permintaan HTTP.
	// Dalam contoh ini, token CSRF akan dicari dalam data formulir dengan nama "_csrf".
	KeyLookup: "form:_csrf",
	// Nama cookie untuk menyimpan token CSRF.
	CookieName: "csrf_",
	// SameSite menentukan kebijakan SameSite untuk cookie CSRF.
	CookieSameSite: "Strict",
	// Durasi berapa lama token CSRF akan berlaku sebelum kedaluwarsa.
	Expiration: 1 * time.Hour,
	// KeyGenerator adalah fungsi untuk menghasilkan token CSRF baru.
	KeyGenerator: GenerateSha256Token,
	// KeyGenerator: utils.UUIDv4,
	// ContextKey adalah kunci konteks di mana token CSRF akan disimpan.
	ContextKey: "csrf_token",
	// CookieSecure menentukan apakah cookie CSRF hanya akan disertakan dalam koneksi HTTPS.
	CookieSecure: true,
	// CookieSessionOnly menentukan apakah cookie CSRF hanya akan disertakan dalam sesi browser.
	CookieSessionOnly: true,
	// CookieHTTPOnly menentukan apakah cookie CSRF hanya dapat diakses melalui HTTP.
	CookieHTTPOnly: false,
	// ErrorHandler adalah fungsi penanganan kesalahan yang akan dipanggil jika token CSRF tidak valid atau hilang.
	ErrorHandler: CsrfErrorHandler,
	// Session adalah penyimpanan sesi yang akan digunakan untuk menyimpan token CSRF.
	Session: session.New(ConfigSession),
	// SessionKey adalah kunci untuk menyimpan token CSRF dalam sesi.
	SessionKey: "legit.csrf.token",
	// HandlerContextKey adalah kunci konteks di mana handler CSRF akan disimpan.
	HandlerContextKey: "legit.csrf.handler",
}

// CSRF Error handler
func CsrfErrorHandler(c *fiber.Ctx, err error) error {
	llog := legitConfig.InitLogger("logs")
	llog.Errorf("CSRF Error: %v Request: %v From: %v", err, c.OriginalURL(), c.IP())
	switch c.Accepts("html", "json") {
	case "json":
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": "403 Forbidden - CSRF Token Not Valid!",
		})
	case "html":
		return c.Status(fiber.StatusForbidden).Render("errors/errors", fiber.Map{
			"Title":   "403 Forbidden - CSRF Token Not Valid!",
			"Code":    "403",
			"Message": "CSRF Token Not Valid!",
		})
	default:
		return c.Status(fiber.StatusForbidden).SendString("403 Forbidden")
	}
}

func GenerateSha256Token() string {
	token := make([]byte, 64)
	rand.Read(token)
	hash := sha256.Sum256(token)
	encodedToken := base64.URLEncoding.EncodeToString(hash[:])
	return encodedToken
}
