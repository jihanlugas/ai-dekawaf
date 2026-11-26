package securityprofile

import (
	"ai-dekawaf/app/base"
	"ai-dekawaf/model"
	"ai-dekawaf/request"
	"strings"

	"gorm.io/gorm"
)

type Repository interface {
	Page(req request.PageSecurityProfile) (mSecurityProfiles []model.SecurityProfile, count int64, err error)
	GetByID(id string) (model.SecurityProfile, error)
	Create(data model.SecurityProfile) error
	Save(data model.SecurityProfile) error
	Updates(data model.SecurityProfile) error
	Delete(data model.SecurityProfile) error
	DeleteByID(id string) error
}

type repository struct {
	db *gorm.DB
}

func (r *repository) GetByID(id string) (model.SecurityProfile, error) {
	var securityProfile model.SecurityProfile
	err := r.db.First(&securityProfile, "id = ?", id).Error
	return securityProfile, err
}

func (r *repository) Create(data model.SecurityProfile) error {
	return r.db.Create(&data).Error
}

func (r *repository) Page(req request.PageSecurityProfile) (mSecurityProfiles []model.SecurityProfile, count int64, err error) {
	query := r.db.Model(&model.SecurityProfile{})
	if req.ProjectID != "" {
		query = query.Where("project_id = ?", req.ProjectID)
	}
	if req.Name != "" {
		query = query.Where("name ILIKE ?", "%"+req.Name+"%")
	}
	if req.Description != "" {
		query = query.Where("description ILIKE ?", "%"+req.Description+"%")
	}
	if req.Domain != "" {
		query = query.Where("domain ILIKE ?", "%"+req.Domain+"%")
	}
	if req.Node != "" {
		query = query.Where("node ILIKE ?", "%"+req.Node+"%")
	}
	if req.Pool != "" {
		query = query.Where("pool ILIKE ?", "%"+req.Pool+"%")
	}
	if req.Port != "" {
		query = query.Where("port = ?", req.Port)
	}
	if req.Destination != "" {
		query = query.Where("destination ILIKE ?", "%"+req.Destination+"%")
	}
	if req.EnforcementMode != "" {
		query = query.Where("enforcement_mode ILIKE ?", "%"+req.EnforcementMode+"%")
	}
	if req.BotProtection != nil {
		query = query.Where("bot_protection = ?", *req.BotProtection)
	}

	query = base.ApplyGlobalSearch(query, req.Search, req)

	err = query.Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	sortField := "created_at"
	sortOrder := "desc"

	if req.SortField != "" {
		sortField = req.SortField
	}

	if strings.EqualFold(req.SortOrder, "asc") {
		sortOrder = "asc"
	} else if strings.EqualFold(req.SortOrder, "desc") {
		sortOrder = "desc"
	}

	query = query.Order(sortField + " " + sortOrder)

	query = query.Offset((req.Page - 1) * req.GetLimit()).Limit(req.GetLimit())
	err = query.Find(&mSecurityProfiles).Error
	if err != nil {
		return nil, 0, err
	}

	return mSecurityProfiles, count, err
}

func (r *repository) Delete(data model.SecurityProfile) error {
	return r.db.Delete(&data).Error
}

func (r *repository) DeleteByID(id string) error {
	return r.db.Delete(&model.SecurityProfile{}, "id = ?", id).Error
}

func (r *repository) Updates(data model.SecurityProfile) error {
	return r.db.Updates(&data).Error
}

func (r *repository) Save(data model.SecurityProfile) error {
	return r.db.Save(&data).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}
