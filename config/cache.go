package config

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/storage/sqlite3"
)

// Inisialisasi penyimpanan cache
var ConfigCache = cache.Config{
	// Middleware selanjutnya dalam rantai middleware, nil berarti tidak ada middleware selanjutnya setelah cache.
	Next: nil,
	// Waktu kadaluarsa untuk setiap item dalam cache.
	Expiration: 30 * time.Minute,
	// Header yang akan digunakan untuk menandai respons yang di-cache.
	CacheHeader: "X-Cache",
	// Mengontrol apakah header Cache-Control akan ditambahkan ke respons.
	CacheControl: true,
	// Fungsi untuk menghasilkan kunci unik berdasarkan konteks permintaan.
	KeyGenerator: func(c *fiber.Ctx) string {
		return GenerateSha256Token()
	},
	// Fungsi untuk menghasilkan waktu kadaluarsa khusus untuk setiap item dalam cache (jika nil, maka menggunakan Expiration).
	ExpirationGenerator: nil,
	// Menyimpan header respons ke dalam cache.
	StoreResponseHeaders: false,
	// Penyimpanan cache yang akan digunakan.
	Storage: sqlite3.New(),
	// Tentukan ukuran maksimum cache jika diperlukan
	// Isi dengan 0 untuk unlimited
	MaxBytes: 100 * 1024 * 1024, // Contoh: 100 MB
	// Metode HTTP yang akan di-cache.
	Methods: []string{fiber.MethodGet, fiber.MethodHead},
}
