package repositories_test

import (
	"MentalHealthCare/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	t.Run("GetAll | Valid", func(t *testing.T) {
		recommendationRepository.On("GetAll").Return([]models.Recommendation{}, nil).Once()

		result, err := recommendationService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetAll | Invalid", func(t *testing.T) {
		recommendationRepository.On("GetAll").Return([]models.Recommendation{}, errors.New("error")).Once()

		result, err := recommendationService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		recommendationRepository.On("GetByID", "1").Return(models.Recommendation{}, nil).Once()

		result, err := recommendationService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByID | Invalid", func(t *testing.T) {
		recommendationRepository.On("GetByID", "0").Return(models.Recommendation{}, errors.New("whoops")).Once()

		result, err := recommendationService.GetByID("0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		recommendationRepository.On("Create", models.RecommendationRequest{}).Return(models.Recommendation{}, nil).Once()

		result, err := recommendationService.Create(models.RecommendationRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create | Invalid", func(t *testing.T) {
		recommendationRepository.On("Create", models.RecommendationRequest{}).Return(models.Recommendation{}, errors.New("oops")).Once()

		result, err := recommendationService.Create(models.RecommendationRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		recommendationRepository.On("Update", models.RecommendationRequest{}, "1").Return(models.Recommendation{}, nil).Once()

		result, err := recommendationService.Update(models.RecommendationRequest{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update | Invalid", func(t *testing.T) {
		recommendationRepository.On("Update", models.RecommendationRequest{}, "0").Return(models.Recommendation{}, errors.New("oops")).Once()

		result, err := recommendationService.Update(models.RecommendationRequest{}, "0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		recommendationRepository.On("Delete", "1").Return(nil).Once()

		err := recommendationService.Delete("1")

		assert.Nil(t, err)
	})

	t.Run("Delete | Invalid", func(t *testing.T) {
		recommendationRepository.On("Delete", "0").Return(errors.New("whoops")).Once()

		err := recommendationService.Delete("0")

		assert.NotNil(t, err)
	})
}