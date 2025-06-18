package repository

import (
	"demo/model"

	"gorm.io/gorm"
)

type AdminRepository interface {
	GetAdminByID(id int64) (*model.Admin, error)
	CreateAdmin(admin *model.Admin) error
}

type adminGormRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminGormRepository{db}
}

func (r *adminGormRepository) GetAdminByID(id int64) (*model.Admin, error) {
	var admin model.Admin
	err := r.db.First(&admin, id).Error
	return &admin, err
}

func (r *adminGormRepository) CreateAdmin(admin *model.Admin) error {
	return r.db.Create(admin).Error
}
