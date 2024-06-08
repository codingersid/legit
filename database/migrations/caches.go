package migrations

import (
	"log"

	"gorm.io/gorm"
)

// table caches
func Caches(db *gorm.DB) {
	// versi 1
	if !db.Migrator().HasTable("caches") {
		// Jika tidak ada, maka buat tabel
		if err := db.Exec(`
			CREATE TABLE IF NOT EXISTS caches (
				id BIGINT AUTO_INCREMENT PRIMARY KEY,
				` + "`key`" + ` VARCHAR(255) UNIQUE,
				value LONGTEXT,
				expires_at DATETIME,
				created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
				updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
			);
		`).Error; err != nil {
			log.Fatal(err)
		}
	}
}
