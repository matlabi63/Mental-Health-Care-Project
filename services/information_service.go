package services

import (
	"MentalHealthCare/models"
	"MentalHealthCare/repositories"
)

type InformationService struct {
	Repository repositories.InformationRepository
}

// InitInformationService initializes the information service.
func InitInformationService() InformationService {
	return InformationService{
		Repository: &repositories.InformationRepositoryImpl{},
	}
}

// GetAll retrieves all information records from the repository.
func (is *InformationService) GetAll() ([]models.Information, error) {
	return is.Repository.GetAll()
}

// GetByID retrieves specific information by ID.
func (is *InformationService) GetByID(id string) (models.Information, error) {
	return is.Repository.GetByID(id)
}

// creates a new information record.
func (is *InformationService) Create(infoReq models.InformationRequest) (models.Information, error) {
	return is.Repository.Create(infoReq)
}

// updates specific information based on the request data and ID.
func (is *InformationService) Update(infoReq models.InformationRequest, id string) (models.Information, error) {
	return is.Repository.Update(infoReq, id)
}

// removes an information record by its ID.
func (is *InformationService) Delete(id string) error {
	return is.Repository.Delete(id)
}
