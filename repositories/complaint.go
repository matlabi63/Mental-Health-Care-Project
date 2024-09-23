package repositories

import (
	"MentalHealthCare/database"
	"MentalHealthCare/models"
)

type ComplaintRepositoryImpl struct{}

func InitComplaintRepository() ComplaintRepository {
	return &ComplaintRepositoryImpl{}
}

func (cr *ComplaintRepositoryImpl) GetAll() ([]models.Complaint, error) {
	var complaints []models.Complaint

	if err := database.DB.Find(&complaints).Error; err != nil {
		return nil, err
	}

	return complaints, nil
}

func (cr *ComplaintRepositoryImpl) GetByID(id string) (models.Complaint, error) {
	var complaint models.Complaint

	if err := database.DB.First(&complaint, "id = ?", id).Error; err != nil {
		return models.Complaint{}, err
	}

	return complaint, nil
}

func (cr *ComplaintRepositoryImpl) Create(complaintReq models.Complaint) (models.Complaint, error) {
	result := database.DB.Create(&complaintReq)

	if err := result.Error; err != nil {
		return models.Complaint{}, err
	}

	return complaintReq, nil
}

func (cr *ComplaintRepositoryImpl) Update(complaintReq models.Complaint, id string) (models.Complaint, error) {
	complaint, err := cr.GetByID(id)

	if err != nil {
		return models.Complaint{}, err
	}

	complaint.UserID = complaintReq.UserID
	complaint.Details = complaintReq.Details

	if err := database.DB.Save(&complaint).Error; err != nil {
		return models.Complaint{}, err
	}

	return complaint, nil
}

func (cr *ComplaintRepositoryImpl) Delete(id string) error {
	complaint, err := cr.GetByID(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&complaint).Error; err != nil {
		return err
	}

	return nil
}
