package repositories_test

import (
	"MentalHealthCare/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllInformation(t *testing.T) {
	t.Run("GetAll | Valid", func(t *testing.T) {
		informationRepository.On("GetAll").Return([]models.Information{}, nil).Once()

		result, err := informationService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetAll | Invalid", func(t *testing.T) {
		informationRepository.On("GetAll").Return([]models.Information{}, errors.New("error")).Once()

		result, err := informationService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByIDInformation(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		informationRepository.On("GetByID", "1").Return(models.Information{}, nil).Once()

		result, err := informationService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByID | Invalid", func(t *testing.T) {
		informationRepository.On("GetByID", "0").Return(models.Information{}, errors.New("whoops")).Once()

		result, err := informationService.GetByID("0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreateInformation(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		informationRepository.On("Create", models.InformationRequest{}).Return(models.Information{}, nil).Once()

		result, err := informationService.Create(models.InformationRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)

		informationRepository.AssertExpectations(t) // new 
	})

	t.Run("Create | Invalid", func(t *testing.T) {
		informationRepository.On("Create", models.InformationRequest{}).Return(models.Information{}, errors.New("oops")).Once()

		result, err := informationService.Create(models.InformationRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
		assert.EqualError(t, err, "oops") // new include

		informationRepository.AssertExpectations(t) // new include
	})
}

func TestUpdateInformation(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		informationRepository.On("Update", models.InformationRequest{}, "1").Return(models.Information{}, nil).Once()

		result, err := informationService.Update(models.InformationRequest{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update | Invalid", func(t *testing.T) {
		informationRepository.On("Update", models.InformationRequest{}, "0").Return(models.Information{}, errors.New("oops")).Once()

		result, err := informationService.Update(models.InformationRequest{}, "0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDeleteInformation(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		informationRepository.On("Delete", "1").Return(nil).Once()

		err := informationService.Delete("1")

		assert.Nil(t, err)
	})

	t.Run("Delete | Invalid", func(t *testing.T) {
		informationRepository.On("Delete", "0").Return(errors.New("whoops")).Once()

		err := informationService.Delete("0")

		assert.NotNil(t, err)
	})
}