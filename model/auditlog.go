package model

import (
	"ai-dekawaf/helper"
	"time"

	"gorm.io/gorm"
)

// Auditlog ..
type Auditlog struct {
	ID             string         `gorm:"primaryKey;type:uuid" json:"id"`
	UserID         string         `gorm:"not null;type:uuid;index" json:"user_id,omitempty"`
	OrganizationID string         `gorm:"not null;type:uuid;index" json:"organization_id,omitempty"`
	Organization   Organization   `gorm:"foreignKey:OrganizationID" json:"organization"`
	ProjectID      string         `gorm:"not null;type:uuid;index" json:"project_id,omitempty"`
	Project        Project        `gorm:"foreignKey:ProjectID" json:"project"`
	IPAddress      string         `gorm:"not null" json:"ip_address"`
	ServiceID      string         `gorm:"not null" json:"service_id,omitempty"`
	ServiceName    string         `gorm:"not null" json:"service_name"`
	ServiceType    string         `gorm:"not null" json:"service_type"`
	MethodName     string         `gorm:"not null" json:"method_name"`
	Action         string         `gorm:"not null" json:"action"`
	Level          string         `gorm:"not null" json:"level"`
	Metadata       string         `gorm:"not null" json:"metadata"`
	CreatedAt      time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	User           *User          `gorm:"foreignKey:UserID" json:"user"`
}

func (auditlog *Auditlog) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	if auditlog.ID == "" {
		auditlog.ID = helper.GetUniqueID()
	}
	if auditlog.CreatedAt.IsZero() {
		auditlog.CreatedAt = now
	}
	if auditlog.UpdatedAt.IsZero() {
		auditlog.UpdatedAt = now
	}
	return nil
}

func (auditlog *Auditlog) BeforeUpdate(tx *gorm.DB) (err error) {
	auditlog.UpdatedAt = time.Now()
	return nil
}

func (Auditlog) TableName() string {
	return "public.auditlogs"
}
