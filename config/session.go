package config

import (
	"encoding/json"
	"time"

	legitConfig "github.com/codingersid/legit-cli/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/utils/v2"
	"gorm.io/gorm"
)

// struct ke dan dari database
type Sessions struct {
	ID        string    `gorm:"primaryKey"`
	Data      string    `gorm:"type:json"`
	ExpiresAt time.Time `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// GormSessionStore adalah struct untuk menyimpan sesi menggunakan Gorm
type GormSessionStore struct {
	db *gorm.DB
}

// NewGormSessionStore membuat instance baru dari GormSessionStore
func NewGormSessionStore(db *gorm.DB) *GormSessionStore {
	return &GormSessionStore{db: db}
}

// Set menyimpan sesi ke dalam database
func (store *GormSessionStore) Set(id string, data []byte, exp time.Duration) error {
	encodedData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	var session Sessions
	// Cek apakah sesi sudah ada di database
	if err := store.db.First(&session, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Jika sesi tidak ditemukan, buat sesi baru
			now := time.Now()
			session = Sessions{
				ID:        id,
				Data:      string(encodedData),
				ExpiresAt: now.Add(exp),
				CreatedAt: now, // Set CreatedAt here
				UpdatedAt: now, // Set UpdatedAt to current time
			}
			return store.db.Create(&session).Error
		}
		return err
	}
	// Jika sesi ditemukan, perbarui data saja
	session.Data = string(encodedData)
	session.UpdatedAt = time.Now()
	return store.db.Save(&session).Error
}

// Get mengambil sesi dari database
func (store *GormSessionStore) Get(id string) ([]byte, error) {
	var session Sessions
	if err := store.db.First(&session, "id = ? AND expires_at > ?", id, time.Now()).Error; err != nil {
		return nil, err
	}
	var decodedData []byte
	if err := json.Unmarshal([]byte(session.Data), &decodedData); err != nil {
		return nil, err
	}
	return decodedData, nil
}

// Delete menghapus sesi dari database
func (store *GormSessionStore) Delete(id string) error {
	return store.db.Delete(&Sessions{}, "id = ?", id).Error
}

// Reset clears all sessions from the database
func (store *GormSessionStore) Reset() error {
	return store.db.Exec("DELETE FROM sessions").Error
}

// Close menutup koneksi database (tidak diperlukan untuk Gorm, tapi untuk memenuhi antarmuka fiber.Storage)
func (store *GormSessionStore) Close() error {
	sqlDB, err := store.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

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
	Storage: getStorageSession(),
}

func getStorageSession() fiber.Storage {
	env := legitConfig.LoadEnv()
	if env["APP_NO_DB"] == "false" {
		sessionInitDb := InitDBWithoutError()
		sessionStore := NewGormSessionStore(sessionInitDb)
		return sessionStore
	}
	return nil
}
