package config

import (
	"encoding/base64"
	"log"
	"os"
	"time"

	legitConfig "github.com/codingersid/legit-cli/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"gorm.io/gorm"
)

func init() {
	// Membuka file log atau membuatnya jika belum ada
	logFile, err := os.OpenFile("logs/logs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Gagal membuka atau membuat file log: %v", err)
	}

	// Mengatur logger untuk menulis ke file
	log.SetOutput(logFile)
}

type Caches struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Key       string    `gorm:"type:varchar(255);uniqueIndex"`
	Value     string    `gorm:"type:longtext"`
	ExpiresAt time.Time `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GormCacheStore struct {
	DB *gorm.DB
}

func (store *GormCacheStore) Get(key string) ([]byte, error) {
	var cache Caches
	err := store.DB.Where("`key` = ? AND `expires_at` > ?", key, time.Now()).First(&cache).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Printf("Kesalahan saat mengambil cache untuk kunci %s: %v", key, err)
		}
		return nil, nil
	}
	decodedValue, err := base64.StdEncoding.DecodeString(cache.Value)
	if err != nil {
		log.Printf("Kesalahan saat mendekode nilai cache untuk kunci %s: %v", key, err)
		return nil, nil
	}

	return decodedValue, nil
}

func (store *GormCacheStore) Set(key string, val []byte, ttl time.Duration) error {
	var cache Caches

	// Meng-encode nilai ke base64
	encodedValue := base64.StdEncoding.EncodeToString(val)
	expiresAt := time.Now().Add(ttl)

	// Memeriksa apakah cache sudah ada berdasarkan kunci
	err := store.DB.Where("`key` = ?", key).First(&cache).Error

	// Jika cache ada, perbarui
	if err == nil { // Tidak ada kesalahan, cache ditemukan
		cache.Value = encodedValue
		cache.ExpiresAt = expiresAt
	} else if err == gorm.ErrRecordNotFound { // Cache tidak ditemukan, buat baru
		cache = Caches{
			Key:       key,
			Value:     encodedValue,
			ExpiresAt: expiresAt,
		}
	} else { // Kesalahan lain saat mencari data
		log.Printf("Kesalahan saat mencari cache dengan kunci %s: %v", key, err)
		return nil
	}

	if err := store.DB.Save(&cache).Error; err != nil {
		log.Printf("Kesalahan saat menyimpan cache untuk kunci %s: %v", key, err)
		return nil
	}

	log.Printf("Cache disimpan untuk kunci %s dengan kedaluwarsa %v", key, expiresAt)
	return nil
}

func (store *GormCacheStore) Delete(key string) error {
	err := store.DB.Where("`key` = ?", key).Delete(&Caches{}).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Printf("Kesalahan saat menghapus cache untuk kunci %s: %v", key, err)
		return err
	}
	return nil
}

func (store *GormCacheStore) Clear() error {
	err := store.DB.Delete(&Caches{}, "`expires_at` <= ?", time.Now()).Error
	if err != nil {
		log.Printf("Kesalahan saat membersihkan cache yang kedaluwarsa: %v", err)
		return err
	}
	return nil
}

func (store *GormCacheStore) Close() error {
	sqlDB, err := store.DB.DB()
	if err != nil {
		log.Printf("Kesalahan saat menutup koneksi database: %v", err)
		return err
	}
	return sqlDB.Close()
}

func (store *GormCacheStore) Reset() error {
	err := store.DB.Exec("DELETE FROM caches").Error
	if err != nil {
		log.Printf("Kesalahan saat mereset cache: %v", err)
		return err
	}
	return nil
}

var ConfigCache = cache.Config{
	// Middleware berikutnya dalam rantai, nil berarti tidak ada middleware berikutnya setelah cache.
	Next: nil,
	// Waktu kedaluwarsa untuk setiap item dalam cache.
	Expiration: 30 * time.Minute,
	// Header untuk menandai respons yang di-cache.
	CacheHeader: "X-Cache",
	// Kontrol apakah header Cache-Control akan ditambahkan ke respons.
	CacheControl: true,
	// Metode HTTP yang akan di-cache.
	Methods: []string{fiber.MethodGet, fiber.MethodHead},
	// Fungsi untuk menghasilkan kunci unik berdasarkan konteks permintaan
	KeyGenerator: func(c *fiber.Ctx) string {
		return c.Path()
	},
	// Fungsi untuk menghasilkan waktu kedaluwarsa khusus untuk setiap item dalam cache (jika nil, Expiration digunakan).
	ExpirationGenerator: nil,
	// Menyimpan header respons dalam cache.
	StoreResponseHeaders: false,
	// Penyimpanan cache yang akan digunakan.
	Storage: getStorageCache(),
	// Tentukan ukuran cache maksimum jika diperlukan
	// Tetapkan ke 0 untuk tidak terbatas
	MaxBytes: 100 * 1024 * 1024, // Contoh: 100 MB
}

func getStorageCache() fiber.Storage {
	env := legitConfig.LoadEnv()
	if env["APP_NO_DB"] == "false" {
		cacheInitDb := InitDBWithoutError()
		return &GormCacheStore{DB: cacheInitDb}
	}
	return nil
}
