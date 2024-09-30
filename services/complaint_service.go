package services

import (
	"MentalHealthCare/models"
	"MentalHealthCare/repositories"
)

type ComplaintService struct {
	Repository repositories.ComplaintRepository
}

// Initialize the service
func InitComplaintService() ComplaintService {
	return ComplaintService{
		Repository: &repositories.ComplaintRepositoryImpl{},
	}
}

// get all
func (cs *ComplaintService) GetAll() ([]models.Complaint, error) {
	return cs.Repository.GetAll()
}

// Fetch a complaint by ID
func (cs *ComplaintService) GetByID(id string) (models.Complaint, error) {
	return cs.Repository.GetByID(id)
}

// Create a complaint
func (cs *ComplaintService) Create(complaintReq models.ComplaintRequest) (models.Complaint, error) {
	return cs.Repository.Create(complaintReq)
}

// Update a complaint
func (cs *ComplaintService) Update(complaintReq models.ComplaintRequest,  id string) (models.Complaint, error) {
	return cs.Repository.Update(complaintReq, id)
}

// Delete a complaint
func (cs *ComplaintService) Delete(id string) error {
	return cs.Repository.Delete(id)
}
