package model

import (
	"time"

	"gorm.io/gorm"
)

// Role ..
type Role struct {
	ID        int64          `gorm:"primaryKey"`
	Name      string         `gorm:"not null" json:"name" `
	CreatedAt time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
