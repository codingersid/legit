package migrations

import (
	"log"

	"github.com/codingersid/legit/app/models"
	"gorm.io/gorm"
)

// tabel sessions
func Sessions(db *gorm.DB) {
	if !db.Migrator().HasTable(&models.Sessions{}) {
		if err := db.AutoMigrate(&models.Sessions{}); err != nil {
			log.Fatal(err)
		}
	}
}
