package securityprofile

import (
	"ai-dekawaf/model"

	"gorm.io/gorm"
)

type Repository interface {
	GetByID(id string) (model.SecurityProfile, error)
	Create(data model.SecurityProfile) error
	Update(data model.SecurityProfile) error
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

func (r *repository) Delete(data model.SecurityProfile) error {
	return r.db.Delete(&data).Error
}

func (r *repository) DeleteByID(id string) error {
	return r.db.Delete(&model.SecurityProfile{}, "id = ?", id).Error
}

func (r *repository) Update(data model.SecurityProfile) error {
	return r.db.Save(&data).Error
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db: db,
	}
}
