package response

type UserClaims struct {
	UserID   string `json:"user_id"`
	OrgID    string `json:"organization_id"`
	Role     string `json:"role"`
	UserType string `json:"user_type"`
}

type UserClaimsWithIP struct {
	UserClaims
	IPAddress string `json:"ip_address"`
}
