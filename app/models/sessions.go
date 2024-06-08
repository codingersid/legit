package models

import (
	"time"

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

// BeforeCreate akan dipanggil sebelum membuat data di database
func (s *Sessions) BeforeCreate(db *gorm.DB) (err error) {
	s.CreatedAt = time.Now()
	return nil
}
