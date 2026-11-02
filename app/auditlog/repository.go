package auditlog

import (
	"dekawaf/constant"
	"dekawaf/model"
	"dekawaf/request"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository interface {
	ErrorToLog(userID string, orgID string, projectID string, ipAddr string, serviceID string, serviceName string, serviceType string, methodName string, action string, metadata string) error
	InfoToLog(userID string, orgID string, projectID string, ipAddr string, serviceID string, serviceName string, serviceType string, methodName string, action string, metadata string) error
}

type auditlogRepository struct {
	connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) Repository {
	return &auditlogRepository{connection}
}

func (a *auditlogRepository) ErrorToLog(userID string, orgID string, projectID string, ipAddr string, serviceID string, serviceName string, serviceType string, methodName string, action string, metadata string) error {
	var req = request.Auditlog{
		UserID:         userID,
		OrganizationID: orgID,
		ProjectID:      projectID,
		IPAddress:      ipAddr,
		ServiceID:      serviceID,
		ServiceName:    serviceName,
		ServiceType:    serviceType,
		MethodName:     methodName,
		Action:         action,
		Level:          "ERROR",
		Metadata:       metadata,
		CreatedAt:      time.Now(),
	}

	return a.insertLogToAuditlog(req)
}

func (a *auditlogRepository) InfoToLog(userID string, orgID string, projectID string, ipAddr string, serviceID string, serviceName string, serviceType string, methodName string, action string, metadata string) error {
	var req = request.Auditlog{
		UserID:         userID,
		OrganizationID: orgID,
		ProjectID:      projectID,
		IPAddress:      ipAddr,
		ServiceID:      serviceID,
		ServiceName:    serviceName,
		ServiceType:    serviceType,
		MethodName:     methodName,
		Action:         action,
		Level:          "INFO",
		Metadata:       metadata,
		CreatedAt:      time.Now(),
	}

	return a.insertLogToAuditlog(req)
}

func (a *auditlogRepository) insertLogToAuditlog(req request.Auditlog) error {
	var audit = model.Auditlog{}
	u, err := uuid.NewRandom()
	if err != nil {
		log.Println(fmt.Errorf(constant.FAILED_TO_CREATE_UUID, err))
		return err
	}

	audit.ID = u.String()
	audit.UserID = req.UserID
	audit.OrganizationID = req.OrganizationID
	audit.ProjectID = req.ProjectID
	audit.IPAddress = req.IPAddress
	audit.ServiceID = req.ServiceID
	audit.ServiceName = req.ServiceName
	audit.ServiceType = req.ServiceType
	audit.MethodName = req.MethodName
	audit.Action = req.Action
	audit.Level = req.Level
	audit.Metadata = req.Metadata
	audit.CreatedAt = req.CreatedAt

	if err := a.connection.Create(&audit).Error; err != nil {
		log.Println(err)
	}
	return err
}
