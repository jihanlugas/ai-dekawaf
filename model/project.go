package model

import (
	"dekawaf/constant"
	"errors"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Project ..
type Project struct {
	ID                        string         `gorm:"primaryKey" json:"id"`
	OrganizationID            string         `gorm:"not null;type:uuid" json:"organization_id,omitempty"`
	Name                      string         `gorm:"not null" json:"name"`
	OpenstackProjectID        string         `gorm:"not null" json:"openstack_project_id"`
	DekaprimeProjectID        string         `gorm:"null" json:"dekaprime_project_id"`
	Description               string         `json:"description"`
	Status                    int32          `gorm:"default:0" json:"status"`
	PaymentMethod             string         `gorm:"type:varchar(10);default:'prepaid'" json:"payment_method"`
	PostpaidType              *string        `gorm:"type:varchar(10)" json:"postpaid_type"`
	PostpaidFixedType         *string        `gorm:"null; type:varchar(10)" json:"postpaid_fixed_type"`
	PaymentType               string         `gorm:"null" json:"payment_type"`
	ServiceActivationDate     *time.Time     `json:"service_activation_date"`
	NetworkNumber             *string        `json:"network_number"`
	IntervalBillingPerMonth   int            `gorm:"default: 1" json:"interval_billing_per_month"`
	InvoiceDuePeriodePerDay   int            `gorm:"default: 30" json:"invoice_due_periode_per_day"`
	FixedBillingPricePerMonth int            `gorm:"default: 0" json:"fixed_billing_price_per_month"`
	DocContract               *string        `json:"doc_contract"`
	Region                    string         `json:"region"`
	VclusterName              string         `json:"vcluster_name"`
	VclusterCluster           string         `json:"vcluster_cluster"`
	VatID                     string         `json:"vat_id"`
	VatName                   string         `json:"vat_name"`
	VatAddress                string         `json:"vat_address"`
	VclusterNamespace         string         `json:"vcluster_namespace"`
	ClusterRefNamespace       string         `json:"cluster_ref_namespace"`
	SuspendDate               *time.Time     `gorm:"null" json:"suspend_date"`
	TerminationDate           *time.Time     `gorm:"null" json:"termination_date"`
	CreatedAt                 time.Time      `gorm:"null" json:"created_at"`
	UpdatedAt                 time.Time      `gorm:"null" json:"updated_at"`
	DeletedAt                 gorm.DeletedAt `gorm:"index" json:",omitempty"`
	VclusterTemplate          string         `json:"vcluster_template"`
	IsPolicyProjectCreated    bool           `json:"is_policy_project_created"`
	IsSubnetProjectCreated    bool           `json:"is_subnet_project_created"`
	IsCiliumProjectCreated    bool           `json:"is_cilium_project_created"`
	IsCrd                     bool           `json:"is_crd"`
	IsDekaRespatiEnable       bool           `json:"is_deka_respati_enable" gorm:"default:false"`
	IsRafayProjectCreated     bool           `gorm:"not null;default:false" json:"is_rafay_project_created"`

	Organization *Organization `gorm:"foreignKey:OrganizationID" json:"organization"`
}

type ProjectFilter struct {
	Id                 *uuid.UUID
	NotInId            []uuid.UUID
	OrgId              *uuid.UUID
	WithDeleted        bool
	HaveNoServiceLimit *bool
}

func (p *Project) IsValidActiveProject(comparisonDate *time.Time) (bool, error) {
	currentTime := time.Now()
	if comparisonDate != nil {
		currentTime = *comparisonDate
	}

	if p.DeletedAt.Valid {
		return false, errors.New("project has been deleted")
	}

	if p.Status != 1 || p.ServiceActivationDate == nil || currentTime.Before(*p.ServiceActivationDate) {
		return false, errors.New("project is not active")
	}

	return true, nil
}

func (p *Project) IsProjectPrepaid() bool {
	return p.PaymentMethod == string(constant.PaymentTypePrepaid)
}

func (p *Project) IsProjectPostpaid() bool {
	return p.PaymentMethod == string(constant.PaymentTypePostpaid)
}
