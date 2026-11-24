package model

import (
	"ai-dekawaf/helper"
	"time"

	"gorm.io/gorm"
)

type SecurityProfile struct {
	ID              string         `gorm:"primaryKey"`
	Name            string         `gorm:"not null"`
	Description     string         `gorm:"not null"`
	Domain          string         `gorm:"not null"`
	Node            string         `gorm:"not null"`
	Pool            string         `gorm:"not null"`
	VirtualServer   string         `gorm:"not null"`
	Destination     string         `gorm:"not null"`
	Policy          string         `gorm:"not null"`
	ClientSSL       string         `gorm:"not null"`
	EnforcementMode string         `gorm:"not null"`
	BotProtection   bool           `gorm:"not null"`
	CreatedAt       time.Time      `gorm:"not null"`
	UpdatedAt       time.Time      `gorm:"not null"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}

func (securityProfile *SecurityProfile) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	if securityProfile.ID == "" {
		securityProfile.ID = helper.GetUniqueID()
	}
	if securityProfile.CreatedAt.IsZero() {
		securityProfile.CreatedAt = now
	}
	if securityProfile.UpdatedAt.IsZero() {
		securityProfile.UpdatedAt = now
	}
	return nil
}

func (securityProfile *SecurityProfile) BeforeUpdate(tx *gorm.DB) (err error) {
	securityProfile.UpdatedAt = time.Now()
	return nil
}

func (SecurityProfile) TableName() string {
	return "dekawaf.security_profiles"
}
