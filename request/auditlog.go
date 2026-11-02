package request

import "time"

type Auditlog struct {
	UserID         string    `json:"user_id"`
	OrganizationID string    `json:"organization_id"`
	ProjectID      string    `json:"project_id"`
	IPAddress      string    `json:"ip_address"`
	ServiceID      string    `json:"service_id"`
	ServiceName    string    `json:"service_name"`
	ServiceType    string    `json:"service_type"`
	MethodName     string    `json:"method_name"`
	Action         string    `json:"action"`
	Level          string    `json:"level"`
	Metadata       string    `json:"metadata"`
	CreatedAt      time.Time `json:"created_at"`
}
