package model

import (
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

type OrganizationActivities struct {
	Id             string    `gorm:"primaryKey;type:uuid" json:"id"`
	OrganizationId string    `gorm:"not null" json:"organization_id,omitempty"`
	UserID         string    `gorm:"null;type:uuid;index" json:"user_id" `
	Action         string    `gorm:"not null" json:"action" `
	EventName      string    `json:"event_name"`
	Description    string    `json:"description"  `
	CreatedAt      time.Time `gorm:"not null" json:"created_at"`
	MetaData       string    `json:"meta_data"`
	User           User      `gorm:"foreignKey:UserID" json:"user"`
}
