package config

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/sqlite3"
	"github.com/gofiber/utils/v2"
)

type Source string

const (
	SourceCookie   Source = "cookie"
	SourceHeader   Source = "header"
	SourceURLQuery Source = "query"
	sessionName    Source = "session_id"
)

// Konfigurasi untuk middleware session
var ConfigSession = session.Config{
	// Expiration menentukan berapa lama sesi akan berlangsung sebelum kedaluwarsa.
	Expiration: 24 * time.Hour,
	// KeyLookup menentukan lokasi di mana ID sesi akan dicari dalam permintaan HTTP.
	// Dalam contoh ini, ID sesi akan dicari dalam cookie dengan nama "session_id".
	KeyLookup: "cookie:session_id",
	// CookieDomain adalah domain yang terkait dengan cookie sesi.
	CookieDomain: "",
	// CookiePath adalah jalur yang terkait dengan cookie sesi.
	CookiePath: "",
	// CookieSecure menentukan apakah cookie sesi hanya akan disertakan dalam koneksi HTTPS.
	CookieSecure: true,
	// CookieHTTPOnly menentukan apakah cookie sesi hanya dapat diakses melalui HTTP.
	CookieHTTPOnly: false,
	// CookieSameSite menentukan kebijakan SameSite untuk cookie sesi.
	CookieSameSite: "Strict",
	// CookieSessionOnly menentukan apakah cookie sesi hanya akan disertakan dalam sesi browser.
	CookieSessionOnly: true,
	// KeyGenerator adalah fungsi untuk menghasilkan ID sesi baru.
	KeyGenerator: utils.UUIDv4,
	// Storage adalah penyimpanan yang akan digunakan untuk menyimpan data sesi.
	// Dalam contoh ini, menggunakan penyimpanan SQLite.
	Storage: sqlite3.New(),
}
