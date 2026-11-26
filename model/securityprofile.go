package model

import (
	"ai-dekawaf/helper"
	"time"

	"gorm.io/gorm"
)

type SecurityProfileStatus string

const (
	SecurityProfileStatusActive   SecurityProfileStatus = "active"
	SecurityProfileStatusPending  SecurityProfileStatus = "pending"
	SecurityProfileStatusInactive SecurityProfileStatus = "inactive"
)

type SecurityProfile struct {
	ID              string                `gorm:"primaryKey" json:"id"`
	OrganizationID  string                `gorm:"not null" json:"organizationId"`
	ProjectID       string                `gorm:"not null" json:"projectId"`
	Name            string                `gorm:"not null" json:"name"`
	Description     string                `gorm:"not null" json:"description"`
	Domain          string                `gorm:"not null" json:"domain"`
	Node            string                `gorm:"not null" json:"node"`
	Pool            string                `gorm:"not null" json:"pool"`
	CertificatePath string                `gorm:"not null" json:"certificatePath"`
	ChainPath       string                `gorm:"not null" json:"chainPath"`
	KeyPath         string                `gorm:"not null" json:"keyPath"`
	ClientSSL       string                `gorm:"not null" json:"clientSsl"`
	VirtualServer   string                `gorm:"not null" json:"virtualServer"`
	Destination     string                `gorm:"not null" json:"destination"`
	Policy          string                `gorm:"not null" json:"policy"`
	EnforcementMode string                `gorm:"not null" json:"enforcementMode"`
	BotProtection   bool                  `gorm:"not null" json:"botProtection"`
	Status          SecurityProfileStatus `gorm:"not null" json:"status"`
	CreatedAt       time.Time             `gorm:"not null" json:"createdAt"`
	UpdatedAt       time.Time             `gorm:"not null" json:"updatedAt"`
	DeletedAt       gorm.DeletedAt        `gorm:"index" json:"-"`
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
