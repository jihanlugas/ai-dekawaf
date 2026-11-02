package project

import (
	"dekawaf/logger"
	"dekawaf/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	GetProjectByUserIdAndProjectId(projectId, userId string) (userProject model.Userproject, err error)
}

type projectRepository struct {
	connection *gorm.DB
}

func NewProjectRepository(connection *gorm.DB) Repository {
	return &projectRepository{connection}
}

func (db *projectRepository) GetProjectByUserIdAndProjectId(projectId, userId string) (model.Userproject, error) {
	var userProject model.Userproject
	err := db.connection.
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
