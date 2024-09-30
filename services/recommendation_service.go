package services

import (
	"MentalHealthCare/models"
	"MentalHealthCare/repositories"
)

type RecommendationService struct {
	Repository repositories.RecommendationRepository
}

// InitRecommendationService initializes the recommendation service.
func InitRecommendationService() RecommendationService {
	return RecommendationService{
		Repository: &repositories.RecommendationRepositoryImpl{},
	}
}

// GetAll retrieves all recommendation records from the repository.
func (rs *RecommendationService) GetAll() ([]models.Recommendation, error) {
	return rs.Repository.GetAll()
}

// GetByID retrieves specific recommendation by ID.
func (rs *RecommendationService) GetByID(id string) (models.Recommendation, error) {
	return rs.Repository.GetByID(id)
}

// creates a new recommendation record.
func (rs *RecommendationService) Create(recommendationReq models.RecommendationRequest) (models.Recommendation, error) {
	return rs.Repository.Create(recommendationReq)
}

// updates specific recommendation based on the request data and ID.
func (rs *RecommendationService) Update(recommendationReq models.RecommendationRequest, id string) (models.Recommendation, error) {
	return rs.Repository.Update(recommendationReq, id)
}

// removes a recommendation record by its ID.
func (rs *RecommendationService) Delete(id string) error {
	return rs.Repository.Delete(id)
}
