package models

import (
	"time"

	"gorm.io/gorm"
)

// struct ke dan dari database
type Caches struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Key       string    `gorm:"type:varchar(255);uniqueIndex"`
	Value     string    `gorm:"type:longtext"`
	ExpiresAt time.Time `gorm:"index"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// BeforeCreate akan dipanggil sebelum membuat data di database
func (s *Caches) BeforeCreate(db *gorm.DB) (err error) {
	s.CreatedAt = time.Now()
	return nil
}
