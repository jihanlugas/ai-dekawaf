package project

import (
	"ai-dekawaf/logger"
	"ai-dekawaf/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	GetProjectByUserIdAndProjectId(projectId, userId string) (userProject model.Userproject, err error)
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
