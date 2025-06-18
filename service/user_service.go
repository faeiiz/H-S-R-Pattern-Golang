package service

import (
	"demo/dto"
	"demo/model"
	"demo/repository"
)

type AdminService interface {
	CreateAdmin(req dto.CreateAdminRequest) (*dto.AdminResponse, error)
	GetAdmin(id int64) (*dto.AdminResponse, error)
}
type adminService struct {
	repo repository.AdminRepository
}

func NewAdminService(repo repository.AdminRepository) AdminService {
	return &adminService{repo}
}

func (s *adminService) CreateAdmin(req dto.CreateAdminRequest) (*dto.AdminResponse, error) {
	admin := &model.Admin{
		Username: req.Name,
		Email:    req.Email,
		Password: req.Password, // In real apps, hash the password!
	}
	err := s.repo.CreateAdmin(admin)
	if err != nil {
		return nil, err
	}
	return &dto.AdminResponse{ID: admin.ID, Name: admin.Username, Email: admin.Email}, nil
}

func (s *adminService) GetAdmin(id int64) (*dto.AdminResponse, error) {
	admin, err := s.repo.GetAdminByID(id)
	if err != nil {
		return nil, err
	}
	return &dto.AdminResponse{ID: admin.ID, Name: admin.Username, Email: admin.Email}, nil
}
