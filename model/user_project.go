package model

import (
	"time"

	"gorm.io/gorm"
)

// Userproject model for user project.
type Userproject struct {
	ID             string         `gorm:"primaryKey;type:uuid" json:"id"`
	UserID         string         `gorm:"not null;type:uuid" json:"user_id,omitempty"`
	ProjectID      string         `gorm:"not null" json:"project_id,omitempty"`
	Project        Project        `gorm:"foreignkey:ProjectID"`
	User           User           `gorm:"foreignkey:UserID"`
	OrganizationID string         `gorm:"not null;type:uuid" json:"organization_id,omitempty"`
	CreatedAt      time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
