package jwtauth

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt"
)

type Usecase interface {
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	Authorized     bool   `json:"authorized"`
	Cidr           string `json:"cid"`
	Email          string `json:"email"`
	Role           string `json:"role"`
	UserID         string `json:"user_id"`
	OrganizationID string `json:"organization_id"`
	UserType       string `json:"user_type"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() Usecase {
	return &jwtService{
		issuer:    "secret",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("API_SECRET")
	if secretKey != "" {
		secretKey = "secret"
	}
	return secretKey
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(jwtT *jwt.Token) (any, error) {
		if _, ok := jwtT.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", jwtT.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
}
