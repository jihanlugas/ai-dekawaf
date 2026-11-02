package user

import (
	"dekawaf/logger"
	"dekawaf/model"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Repository interface {
	GetUserByID(string) (model.User, error)
	GetSuperAdminPrivilegesByID(string) ([]model.SuperAdminRoles, error)
	GetAllUser() ([]model.User, error)
}

type userRepository struct {
	connection *gorm.DB
}

func NewUserRepository(connection *gorm.DB) Repository {
	return &userRepository{connection}
}

func (u *userRepository) GetUserByID(id string) (model.User, error) {
	var user model.User
	err := u.connection.
		Preload("Organization").
		Preload("OrganizationRole").
		First(&user, "id = ?", id).Error
	if err != nil {
		logger.Error("GetUserByID", err, logrus.Fields{"user_id": id})
		return user, err
	}
	return user, nil
}

func (u *userRepository) GetSuperAdminPrivilegesByID(id string) ([]model.SuperAdminRoles, error) {
	var superAdminRoles []model.SuperAdminRoles
	err := u.connection.
		Joins("JOIN super_admin_role_mapping srm ON srm.id_role = super_admin_roles.id").
		Where("srm.id_admin = ?", id).
		Find(&superAdminRoles).Error
	if err != nil {
		logger.Error("GetSuperAdminPrivilegesByID", err, logrus.Fields{"user_id": id})
		return superAdminRoles, err
	}
	return superAdminRoles, nil
}

func (u *userRepository) GetAllUser() ([]model.User, error) {
	var (
		users []model.User
	)
	if err := u.connection.Find(&users).Error; err != nil {
		return users, nil
	}

	return users, nil
}
