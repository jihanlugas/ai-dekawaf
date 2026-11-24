package user

import (
	"ai-dekawaf/logger"
	"ai-dekawaf/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	GetUserByID(id string, preload ...string) (model.User, error)
	GetSuperAdminPrivilegesByID(string) ([]model.SuperAdminRoles, error)
	GetAllUser() ([]model.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) GetUserByID(id string, preload ...string) (model.User, error) {
	var user model.User
	conn := r.db
	for _, p := range preload {
		conn = conn.Preload(p)
	}
	err := conn.
		First(&user, "id = ?", id).Error
	if err != nil {
		logger.Error("GetUserByID", err, logrus.Fields{"user_id": id})
		return user, err
	}
	return user, nil
}

func (r *repository) GetSuperAdminPrivilegesByID(id string) ([]model.SuperAdminRoles, error) {
	var superAdminRoles []model.SuperAdminRoles
	err := r.db.
		Joins("JOIN super_admin_role_mapping srm ON srm.id_role = super_admin_roles.id").
		Where("srm.id_admin = ?", id).
		Find(&superAdminRoles).Error
	if err != nil {
		logger.Error("GetSuperAdminPrivilegesByID", err, logrus.Fields{"user_id": id})
		return superAdminRoles, err
	}
	return superAdminRoles, nil
}

func (r *repository) GetAllUser() ([]model.User, error) {
	var (
		users []model.User
	)
	if err := r.db.Find(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}
