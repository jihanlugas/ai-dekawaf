package model

import "time"

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
