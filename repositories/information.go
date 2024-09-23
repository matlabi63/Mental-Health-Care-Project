package repositories

import (
	"MentalHealthCare/database"
	"MentalHealthCare/models"
)

type InformationRepositoryImpl struct{}

func InitInformationRepository() InformationRepository {
	return &InformationRepositoryImpl{}
}

func (ir *InformationRepositoryImpl) GetAll() ([]models.Information, error) {
	var informationList []models.Information

	if err := database.DB.Find(&informationList).Error; err != nil {
		return nil, err
	}

	return informationList, nil
}

func (ir *InformationRepositoryImpl) GetByID(id string) (models.Information, error) {
	var information models.Information

	if err := database.DB.First(&information, "id = ?", id).Error; err != nil {
		return models.Information{}, err
	}

	return information, nil
}

func (ir *InformationRepositoryImpl) Create(infoReq models.InformationRequest) (models.Information, error) {
	var information models.Information = models.Information{
		Title: infoReq.Title,
		Content: infoReq.Content,
		Author: infoReq.Author,
	}
	result := database.DB.Create(&information)

	if err := result.Error; err != nil {
		return models.Information{}, err
	}
	if err := result.Last(&information).Error; err != nil {
		return models.Information{}, err
	}

	return information, nil
}

func (ir *InformationRepositoryImpl) Update(infoReq models.InformationRequest, id string) (models.Information, error) {
	information, err := ir.GetByID(id)

	if err != nil {
		return models.Information{}, err
	}

	information.Title = infoReq.Title
	information.Content = infoReq.Content
	information.Author = infoReq.Author

	if err := database.DB.Save(&information).Error; err != nil {
		return models.Information{}, err
	}

	return information, nil
}

func (ir *InformationRepositoryImpl) Delete(id string) error {
	information, err := ir.GetByID(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&information).Error; err != nil {
		return err
	}

	return nil
}
