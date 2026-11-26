package auditlog

import (
	"ai-dekawaf/model"
	"log"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	ErrorToLog(mLog model.Auditlog, action string, metadata string) error
	InfoToLog(mLog model.Auditlog, action string, metadata string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) ErrorToLog(mLog model.Auditlog, action string, metadata string) error {
	mLog.Action = action
	mLog.Level = "ERROR"
	mLog.Metadata = metadata
	mLog.CreatedAt = time.Now()

	return r.insertLogToAuditlog(mLog)
}

func (r *repository) InfoToLog(mLog model.Auditlog, action string, metadata string) error {
	mLog.Action = action
	mLog.Level = "INFO"
	mLog.Metadata = metadata
	mLog.CreatedAt = time.Now()

	return r.insertLogToAuditlog(mLog)
}

func (r *repository) insertLogToAuditlog(mLog model.Auditlog) error {
	if err := r.db.Create(&mLog).Error; err != nil {
		log.Println(err)
	}
	return nil
}
