package services

import (
	"MentalHealthCare/models"
	"MentalHealthCare/repositories"
)

type DoctorService struct {
	Repository repositories.DoctorRepository
}

// InitDoctorService initializes the doctor service.
func InitDoctorService() DoctorService {
	return DoctorService{
		Repository: &repositories.DoctorRepositoryImpl{},
	}
}

// GetAll retrieves all doctors from the repository.
func (ds *DoctorService) GetAll() ([]models.Doctor, error) {
	return ds.Repository.GetAll()
}

// GetByID retrieves a specific doctor by ID.
func (ds *DoctorService) GetByID(id string) (models.Doctor, error) {
	return ds.Repository.GetByID(id)
}

// Create creates a new doctor profile based on the request data.
func (ds *DoctorService) Create(doctorReq models.DoctorRequest) (models.Doctor, error) {
	return ds.Repository.Create(doctorReq)
}

// Update updates a specific doctor profile based on the request data and ID.
func (ds *DoctorService) Update(doctorReq models.DoctorRequest, id string) (models.Doctor, error) {
	return ds.Repository.Update(doctorReq, id)
}

// Delete removes a doctor profile by its ID.
func (ds *DoctorService) Delete(id string) error {
	return ds.Repository.Delete(id)
}
