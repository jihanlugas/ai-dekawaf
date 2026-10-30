package model

import (
	"time"

	"gorm.io/gorm"
)

type AccountBalance struct {
	ID                   string                 `gorm:"primaryKey;type:uuid" json:"id"`
	OrganizationID       string                 `gorm:"type:uuid" json:"organization_id"`
	Balance              float64                `json:"balance"`
	Type                 string                 `gorm:"type:char(1)" json:"type"`
	CreatedAt            time.Time              `gorm:"null" json:"created_at"`
	UpdatedAt            time.Time              `gorm:"null" json:"updated_at"`
	DeletedAt            gorm.DeletedAt         `gorm:"index" json:",omitempty"`
	AccountBalanceDetail []AccountBalanceDetail `gorm:"foreignKey:AccountBalanceID" json:"account_balance_detail"`
	ProjectID            string                 `json:"project_id"`
}

type AccountBalanceDetail struct {
	ID               uint64         `gorm:"primary_key;not null;auto_increment" json:"id"`
	AccountBalanceID string         `gorm:"type:uuid" json:"account_balance_id"`
	Source           string         `gorm:"type:varchar(70)" json:"source"`
	Amount           float64        `json:"amount"`
	Description      string         `json:"description"`
	CurrentBalance   float64        `json:"current_balance"`
	PreviousBalance  float64        `json:"previous_balance"`
	Type             string         `gorm:"type:varchar(7)" json:"type"`
	TransactionID    *int64         `json:"transaction_id"`
	InvoiceID        *string        `json:"invoice_id"`
	CreatedBy        string         `json:"created_by"`
	CreatedAt        time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:",omitempty"`
}
