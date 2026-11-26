package request

import "mime/multipart"

type PageSecurityProfile struct {
	Paging
	ProjectID       string `json:"projectId" form:"projectId" query:"projectId" validate:"required,specialchar"`
	Name            string `json:"name" form:"name" query:"name" search:"name" validate:"specialchar"`
	Description     string `json:"description" form:"description" query:"description" search:"description" validate:"specialchar"`
	Domain          string `json:"domain" form:"domain" query:"domain" search:"domain" validate:"specialchar"`
	Node            string `json:"node" form:"node" query:"node" search:"node" validate:"omitempty,specialchar,lowercase"`
	Pool            string `json:"pool" form:"pool" query:"pool" search:"pool" validate:"omitempty,specialchar,lowercase"`
	Port            string `json:"port" form:"port" query:"port" validate:"omitempty,numeric"`
	Destination     string `json:"destination" form:"destination" query:"destination" search:"destination" validate:"specialchar"`
	EnforcementMode string `json:"enforcementMode" form:"enforcementMode" query:"enforcementMode" search:"enforcement_mode" validate:"omitempty,oneof=transparent blocking"`
	BotProtection   *bool  `json:"botProtection" form:"botProtection" query:"botProtection"`
	Search          string `json:"search" form:"search" query:"search" validate:"specialchar"`
}
type CreateSecurityProfile struct {
	ProjectID         string                `json:"projectId" form:"projectId" validate:"required"`
	Name              string                `json:"name" form:"name" validate:"required,specialchar"`
	Description       string                `json:"description" form:"description" validate:"specialchar"`
	Domain            string                `json:"domain" form:"domain" validate:"required,domain"`
	Node              string                `json:"node" form:"node" validate:"required,specialchar,lowercase"`
	Pool              string                `json:"pool" form:"pool" validate:"required,specialchar,lowercase"`
	Port              string                `json:"port" form:"port" validate:"required,numeric"`
	Destination       string                `json:"destination" form:"destination" validate:"required,ipv4"`
	ServerTecnologies []string              `json:"serverTechnologies" form:"serverTechnologies"`
	EnforcementMode   string                `json:"enforcementMode" form:"enforcementMode" validate:"required,oneof=transparent blocking"`
	BotProtection     bool                  `json:"botProtection" form:"botProtection"`
	CertificateFile   *multipart.FileHeader `json:"certificateFile" form:"certificateFile" validate:"required"`
	KeyFile           *multipart.FileHeader `json:"keyFile" form:"keyFile" validate:"required"`
	ChainFile         *multipart.FileHeader `json:"chainFile" form:"chainFile" validate:""`
}
