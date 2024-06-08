package migrations

import (
	"log"

	"github.com/codingersid/legit/app/models"
	"gorm.io/gorm"
)

// tabel sessions
func Sessions(db *gorm.DB) {
	if !db.Migrator().HasTable(&models.Sessions{}) {
		// Jika tidak ada, maka buat tabel
		if err := db.Exec(`
			CREATE TABLE IF NOT EXISTS sessions (
				id VARCHAR(255) PRIMARY KEY,
				data JSON,
				expires_at DATETIME,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
			);
		`).Error; err != nil {
			log.Fatal(err)
		}
	}
}
