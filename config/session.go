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
	session := Sessions{
		ID:        id,
		Data:      string(encodedData),
		ExpiresAt: time.Now().Add(exp),
	}
	// log.Printf("Setting session: %+v", session)
	return store.db.Save(&session).Error
}

// Get mengambil sesi dari database
func (store *GormSessionStore) Get(id string) ([]byte, error) {
	var session Sessions
	// log.Printf("Getting session with id: %s", id)
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
	// log.Printf("Deleting session with id: %s", id)
	return store.db.Delete(&Sessions{}, "id = ?", id).Error
}

// Reset clears all sessions from the database
func (store *GormSessionStore) Reset() error {
	// log.Println("Resetting all sessions")
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

// Inisialisasi penyimpanan session menggunakan Gorm
var sessionInitDb = InitDB()
var sessionStore = NewGormSessionStore(sessionInitDb)

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
		return sessionStore
	}
	return nil
}
