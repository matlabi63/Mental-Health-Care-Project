package repositories_test

import (
	"MentalHealthCare/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllComplaint(t *testing.T) {
	t.Run("GetAll | Valid", func(t *testing.T) {
		complaintRepository.On("GetAll").Return([]models.Complaint{}, nil).Once()

		result, err := complaintService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetAll | Invalid", func(t *testing.T) {
		complaintRepository.On("GetAll").Return([]models.Complaint{}, errors.New("error")).Once()

		result, err := complaintService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByIDComplaint(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		complaintRepository.On("GetByID", "1").Return(models.Complaint{}, nil).Once()

		result, err := complaintService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByID | Invalid", func(t *testing.T) {
		complaintRepository.On("GetByID", "0").Return(models.Complaint{}, errors.New("oops")).Once()

		result, err := complaintService.GetByID("0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreateComplaint(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		complaintRepository.On("Create", models.ComplaintRequest{}).Return(models.Complaint{}, nil).Once()

		result, err := complaintService.Create(models.ComplaintRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create | Invalid", func(t *testing.T) {
		complaintRepository.On("Create", models.ComplaintRequest{}).Return(models.Complaint{}, errors.New("oops")).Once()

		result, err := complaintService.Create(models.ComplaintRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdateComplaint(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		complaintRepository.On("Update", models.ComplaintRequest{}, "1").Return(models.Complaint{}, nil).Once()

		result, err := complaintService.Update(models.ComplaintRequest{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update | Invalid", func(t *testing.T) {
		complaintRepository.On("Update", models.ComplaintRequest{}, "0").Return(models.Complaint{}, errors.New("whoops")).Once()

		result, err := complaintService.Update(models.ComplaintRequest{}, "0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDeleteComplaint(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		complaintRepository.On("Delete", "1").Return(nil).Once()

		err := complaintService.Delete("1")

		assert.Nil(t, err)
	})

	t.Run("Delete | Invalid", func(t *testing.T) {
		complaintRepository.On("Delete", "0").Return(errors.New("whoops")).Once()

		err := complaintService.Delete("0")

		assert.NotNil(t, err)
	})
}
