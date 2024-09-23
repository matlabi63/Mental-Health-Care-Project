package repositories

import (
	"MentalHealthCare/database"
	"MentalHealthCare/models"
)

type RecommendationRepositoryImpl struct{}

func InitRecommendationRepository() RecommendationRepository {
	return &RecommendationRepositoryImpl{}
}

func (rr *RecommendationRepositoryImpl) GetAll() ([]models.Recommendation, error) {
	var recommendations []models.Recommendation

	if err := database.DB.Find(&recommendations).Error; err != nil {
		return nil, err
	}

	return recommendations, nil
}

func (rr *RecommendationRepositoryImpl) GetByID(id string) (models.Recommendation, error) {
	var recommendation models.Recommendation

	if err := database.DB.First(&recommendation, "id = ?", id).Error; err != nil {
		return models.Recommendation{}, err
	}

	return recommendation, nil
}

func (rr *RecommendationRepositoryImpl) Create(recReq models.RecommendationRequest) (models.Recommendation, error) {
	var recommendation models.Recommendation = models.Recommendation{
		Description: recReq.Description,
		DoctorID: recReq.DoctorID,
		UserID: recReq.UserID,
	}
	result := database.DB.Create(&recommendation)

	if err := result.Error; err != nil {
		return models.Recommendation{}, err
	}

	if err := result.Last(&recommendation).Error; err != nil {
		return models.Recommendation{}, err
	}
	

	return recommendation, nil
}

func (rr *RecommendationRepositoryImpl) Update(recReq models.RecommendationRequest, id string) (models.Recommendation, error) {
	recommendation, err := rr.GetByID(id)

	if err != nil {
		return models.Recommendation{}, err
	}

	recommendation.Description = recReq.Description
	recommendation.DoctorID = recReq.DoctorID
	recommendation.UserID = recReq.UserID

	if err := database.DB.Save(&recommendation).Error; err != nil {
		return models.Recommendation{}, err
	}

	return recommendation, nil
}

func (rr *RecommendationRepositoryImpl) Delete(id string) error {
	recommendation, err := rr.GetByID(id)

	if err != nil {
		return err
	}

	if err := database.DB.Delete(&recommendation).Error; err != nil {
		return err
	}

	return nil
}
