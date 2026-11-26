package project

import (
	"ai-dekawaf/logger"
	"ai-dekawaf/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	GetByID(projectID string, preloads ...string) (model.Project, error)
	GetProjectByUserIdAndProjectId(projectId, userId string) (userProject model.Userproject, err error)
	ProjectAvailabilityStatus(projectID string) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetProjectByUserIdAndProjectId(projectId, userId string) (model.Userproject, error) {
	var userProject model.Userproject
	err := r.db.
		Where(" project_id = ? and user_id = ? ", projectId, userId).
		Find(&userProject).Error
	if err != nil {
		logger.Error("GetProjectByUserIdAndProjectId", err, logrus.Fields{
			"project_id": projectId,
			"user_id":    userId,
		})
		return userProject, err
	}
	return userProject, nil
}

func (r *repository) GetByID(projectID string, preloads ...string) (model.Project, error) {
	var project model.Project
	db := r.db.Where("id = ?", projectID)
	for _, preload := range preloads {
		db = db.Preload(preload)
	}
	err := db.First(&project).Error
	if err != nil {
		logger.Error("GetByID", err, logrus.Fields{
			"project_id": projectID,
		})
		return project, err
	}
	return project, nil
}

func (r *repository) ProjectAvailabilityStatus(projectID string) (bool, error) {
	var count int64
	err := r.db.Model(&model.Project{}).
		Joins("LEFT JOIN organizations on organizations.id::text = projects.organization_id::text").
		Where("projects.id = ?", projectID).
		Where("projects.status != ?", 1).
		Where("(projects.suspend_date IS NOT NULL OR projects.termination_date IS NOT NULL)").
		Where("organizations.suspend_type IS NOT NULL").
		Where("(organizations.suspend_date IS NOT NULL OR organizations.termination_date IS NOT NULL)").
		Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return false, nil
	}

	return true, nil

}
