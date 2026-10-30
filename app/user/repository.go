package user

import (
	"dekawaf/model"

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

// func NewUserRepository(connection *gorm.DB) Repository {
// 	return &userRepository{connection}
// }
