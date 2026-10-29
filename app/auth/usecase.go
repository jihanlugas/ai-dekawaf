package auth

// import (
// 	"dekawaf/constant"
// 	"dekawaf/logger"
// 	"dekawaf/response"
// 	"fmt"
// 	"os"
// 	"strings"

// 	"github.com/gin-gonic/gin"
// 	"github.com/golang-jwt/jwt"
// )

// const (
// 	errFailedParseToken = "Failed parse token: "
// )

// func parseToken(tokenString string) (jwt.MapClaims, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf(constant.ErrUnexpectedSingingMethodErr, token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("API_SECRET")), nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	if !token.Valid {
// 		return nil, fmt.Errorf("invalid token")
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok {
// 		return nil, fmt.Errorf("invalid claims")
// 	}

// 	return claims, nil
// }

// func GetUserID(c *gin.Context) string {
// 	tokenString := ExtractToken(c)

// 	claims, err := parseToken(tokenString)
// 	if err != nil {
// 		logger.Error(errFailedParseToken, err)
// 		return ""
// 	}

// 	userID, _ := claims["user_id"].(string)
// 	return userID
// }

// func GetOrganizationID(c *gin.Context) string {
// 	tokenString := ExtractToken(c)

// 	claims, err := parseToken(tokenString)
// 	if err != nil {
// 		logger.Error(errFailedParseToken, err)
// 		return ""
// 	}

// 	orgID, _ := claims["organization_id"].(string)
// 	return orgID
// }

// func GetRoleAndUserID(c *gin.Context) (string, string) {
// 	tokenString := ExtractToken(c)

// 	claims, err := parseToken(tokenString)
// 	if err != nil {
// 		logger.Error(errFailedParseToken, err)
// 		return "", ""
// 	}

// 	userID, _ := claims["user_id"].(string)
// 	role, _ := claims["role"].(string)

// 	return userID, role
// }

// func GetCurentUser(c *gin.Context) (res response.UserClaims) {
// 	tokenString := ExtractToken(c)

// 	claims, err := parseToken(tokenString)
// 	if err != nil {
// 		logger.Error(errFailedParseToken, err)
// 		return
// 	}

// 	res = response.UserClaims{
// 		UserID:   getStringClaim(claims, "user_id"),
// 		OrgID:    getStringClaim(claims, "organization_id"),
// 		Role:     getStringClaim(claims, "role"),
// 		UserType: getStringClaim(claims, "user_type"),
// 	}

// 	return
// }

// func getStringClaim(claims jwt.MapClaims, key string) string {
// 	if val, ok := claims[key].(string); ok {
// 		return val
// 	}
// 	return ""
// }

// // ExtractToken ..
// func ExtractToken(c *gin.Context) string {
// 	if token := c.Query("token"); token != "" {
// 		return token
// 	}

// 	bearerToken := c.GetHeader("Authorization")
// 	parts := strings.SplitN(bearerToken, " ", 2)
// 	if len(parts) == 2 {
// 		return parts[1]
// 	}

// 	return ""
// }
