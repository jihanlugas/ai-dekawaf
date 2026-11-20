package request

import "mime/multipart"

type CreateSecurityProfile struct {
	Name              string                `json:"name" form:"name" validate:"required,specialchar"`
	Description       string                `json:"description" form:"description" validate:"specialchar"`
	Domain            string                `json:"domain" form:"domain" validate:"required,domain"`
	Node              string                `json:"node" form:"node" validate:"required,specialchar,lowercase"`
	Pool              string                `json:"pool" form:"pool" validate:"required,specialchar,lowercase"`
	Port              string                `json:"port" form:"port" validate:"required,numeric"`
	ServerTecnologies []string              `json:"serverTechnologies" form:"serverTechnologies"`
	EnforcementMode   string                `json:"enforcementMode" form:"enforcementMode" validate:"required,oneof=transparent blocking"`
	BotProtection     bool                  `json:"botProtection" form:"botProtection"`
	CertificateFile   *multipart.FileHeader `json:"certificateFile" form:"certificateFile" validate:"required"`
	KeyFile           *multipart.FileHeader `json:"keyFile" form:"keyFile" validate:"required"`
	ChainFile         *multipart.FileHeader `json:"chainFile" form:"chainFile" validate:""`
}
