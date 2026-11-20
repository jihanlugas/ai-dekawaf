package helper

import (
	"net"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetIP(c *gin.Context) string {
	//Get IP from the X-REAL-IP header
	ip := c.Request.Header.Get("X-Real-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip
	}

	//Get IP from X-FORWARDED-FOR header
	ips := c.Request.Header.Get("X-Forwarded-For")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
	if err != nil {
		return ""
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip
	}
	return ""
}

func GetUniqueID() string {
	uuid := uuid.New().String()
	return uuid
}

func GenerateRandomStringUUID8() string {
	return uuid.New().String()[:8]
}
