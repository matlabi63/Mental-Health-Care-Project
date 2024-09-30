package services

import (
	"MentalHealthCare/models"
	"MentalHealthCare/repositories"
)

type AdminService struct {
    Repository repositories.AdminRepository
}

func InitAdminService() AdminService {
    return AdminService{
        Repository: &repositories.AdminRepositoryImpl{},
    }
}

func (as *AdminService) GetAll() ([]models.Admin, error) {
    return as.Repository.GetAll()
}

func (as *AdminService) GetByID(id string) (models.Admin, error) {
    return as.Repository.GetByID(id)
}

func (as *AdminService) Create(adminReq models.AdminRequest) (models.Admin, error) {
    return as.Repository.Create(adminReq)
}

func (as *AdminService) Update(adminReq models.AdminRequest, id string) (models.Admin, error) {
    return as.Repository.Update(adminReq, id)
}

func (as *AdminService) Delete(id string) error {
    return as.Repository.Delete(id)
}
