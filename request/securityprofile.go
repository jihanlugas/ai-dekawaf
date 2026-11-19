package request

import "mime/multipart"

type CreateSecurityProfile struct {
	Name              string                `json:"name" form:"name"`
	Description       string                `json:"description" form:"description"`
	Domain            string                `json:"domain" form:"domain"`
	Node              string                `json:"node" form:"node"`
	Pool              string                `json:"pool" form:"pool"`
	Port              string                `json:"port" form:"port"`
	ServerTecnologies []string              `json:"serverTechnologies" form:"serverTechnologies"`
	EnforcementMode   string                `json:"enforcementMode" form:"enforcementMode"`
	BotProtection     bool                  `json:"botProtection" form:"botProtection"`
	CertificateFile   *multipart.FileHeader `json:"certificateFile" form:"certificateFile"`
	KeyFile           *multipart.FileHeader `json:"keyFile" form:"keyFile"`
	ChainFile         *multipart.FileHeader `json:"chainFile" form:"chainFile"`
}
