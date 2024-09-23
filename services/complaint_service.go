package services

import (
	"MentalHealthCare/models"
	"MentalHealthCare/repositories"
)

// type ComplaintService struct {
// 	Repository repositories.ComplaintRepository
// }

// // InitComplaintService initializes the complaint service.
// func InitComplaintService() ComplaintService {
// 	return ComplaintService{
// 		Repository: &repositories.ComplaintRepositoryImpl{},
// 	}
// }

// // GetAll retrieves all complaints from the repository.
// func (cs *ComplaintService) GetAll() ([]models.Complaint, error) {
// 	return cs.Repository.GetAll()
// }

// // GetByID retrieves a specific complaint by ID.
// func (cs *ComplaintService) GetByID(id string) (models.Complaint, error) {
// 	return cs.Repository.GetByID(id)
// }

// // Create creates a new complaint based on the request data.
// func (cs *ComplaintService) Create(complaintReq models.ComplaintRequest) (models.Complaint, error) {
//     complaint := complaintReq.toComplaint() // Assuming ToComplaint is implemented
//     return cs.Repository.Create(complaint)
// }

// func (cs *ComplaintService) Update(complaintReq models.ComplaintRequest, id string) (models.Complaint, error) {
// 	complaint := complaintReq.toComplaint() // Use the method here as well
// 	return cs.Repository.Update(complaint, id)
// }

// // Delete removes a complaint by its ID.
// func (cs *ComplaintService) Delete(id string) error {
// 	return cs.Repository.Delete(id)
// }

////


type ComplaintService struct {
	Repository repositories.ComplaintRepository
}

// InitComplaintService initializes the complaint service.
func InitComplaintService() ComplaintService {
	return ComplaintService{
		Repository: &repositories.ComplaintRepositoryImpl{},
	}
}

// GetAll retrieves all complaints from the repository.
func (cs *ComplaintService) GetAll() ([]models.Complaint, error) {
	return cs.Repository.GetAll()
}

// GetByID retrieves a specific complaint by ID.
func (cs *ComplaintService) GetByID(id string) (models.Complaint, error) {
	return cs.Repository.GetByID(id)
}

// Create creates a new complaint based on the request data.
func (cs *ComplaintService) Create(complaintReq models.ComplaintRequest) (models.Complaint, error) {
	// Use the ToComplaint method to convert ComplaintRequest to Complaint
	complaint := complaintReq.ToComplaint()
	return cs.Repository.Create(complaint)
}

// Update updates an existing complaint based on the request data and complaint ID.
func (cs *ComplaintService) Update(complaintReq models.ComplaintRequest, id string) (models.Complaint, error) {
	// Use the ToComplaint method to convert ComplaintRequest to Complaint
	complaint := complaintReq.ToComplaint()
	return cs.Repository.Update(complaint, id)
}

// Delete removes a complaint by its ID.
func (cs *ComplaintService) Delete(id string) error {
	return cs.Repository.Delete(id)
}
