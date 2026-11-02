package middleware

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"dekawaf/constant"
	"dekawaf/logger"
	"dekawaf/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

const (
	noAuthorizationHeaderFound = "No Authorization header found"
	tokenAuthorizationNotValid = "Token authorization not valid"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// AuthorizeJWT validates the token user given, return 401 if not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			err := errors.New("Token not found.")
			logger.Error(noAuthorizationHeaderFound, err)
			response.ErrorAbort(c, http.StatusUnauthorized, noAuthorizationHeaderFound, err, nil)

			return
		}

		typeToken := strings.Split(authHeader, " ")
		if len(typeToken) != 2 || typeToken[0] != "Bearer" {
			logger.Error(tokenAuthorizationNotValid, nil)
			response.ErrorAbort(c, http.StatusUnauthorized, tokenAuthorizationNotValid, nil, nil)

			return
		}

		extractedToken := strings.Split(authHeader, "Bearer ")
		authHeader = strings.TrimSpace(extractedToken[1])
		// logger.Info("Authorization Token", logrus.Fields{"token": authHeader})

		token, err := validateToken(authHeader)
		if err != nil {
			logger.Error(tokenAuthorizationNotValid, nil)
			response.ErrorAbort(c, http.StatusUnauthorized, tokenAuthorizationNotValid, err, nil)

			return
		}

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			// logger.Info("Token is Valid", logrus.Fields{"claims": claims})
			roles := fmt.Sprintf("%s", claims["role"])
			// Bind Token Claims

			// Inject Consumer ID & Type for Prometheus Metrics
			if userID, ok := claims["user_id"].(string); ok && userID != "" {
				c.Writer.Header().Set("X-Consumer-ID", userID)
			}
			if role, ok := claims["role"].(string); ok && (role == "Superadmin" || role == "GodMode") {
				c.Writer.Header().Set("X-Consumer-Type", "admin")
			} else {
				c.Writer.Header().Set("X-Consumer-Type", "customer")
			}

			c.Set(constant.ContextUserRole, roles)
			c.Set(constant.ContextUserID, claims["user_id"])
			c.Set(constant.ContextOrganizationID, claims["organization_id"])
			c.Next()
		} else {
			logger.Error(noAuthorizationHeaderFound, err)
			response.ErrorAbort(c, http.StatusUnauthorized, "Not Valid Token or Token Expired", err, nil)
			return
		}
	}
}

func validateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(jwtT *jwt.Token) (any, error) {
		if _, ok := jwtT.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", jwtT.Header["alg"])
		}
		return []byte(os.Getenv("API_SECRET")), nil
	})
}

// GetOrganizationID returns the organization ID from the gin context.
func GetOrganizationID(c *gin.Context) string {
	return c.GetString(constant.ContextOrganizationID)
}

// GetUser returns the user information from the gin context.
func GetUserID(c *gin.Context) string {
	return c.GetString(constant.ContextUserID)
}

// GetUserRole returns the user role from the gin context.
func GetUserRole(c *gin.Context) string {
	return c.GetString(constant.ContextUserRole)
}

// GetRoleName returns the role name from the role ID.
func GetRoleName(id int64) string {
	var roleName string
	switch id {
	case 1:
		roleName = "Owner"
	case 2:
		roleName = "Member"
	case 3:
		roleName = "Customer-Care"
	case 99:
		roleName = "Superadmin"
	case 100:
		roleName = "GodMode"
	}

	return roleName
}
