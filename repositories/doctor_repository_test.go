package repositories_test

import (
	"MentalHealthCare/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllDoctor(t *testing.T) {
	t.Run("GetAll | Valid", func(t *testing.T) {
		doctorRepository.On("GetAll").Return([]models.Doctor{}, nil).Once()

		result, err := doctorService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetAll | Invalid", func(t *testing.T) {
		doctorRepository.On("GetAll").Return([]models.Doctor{}, errors.New("error")).Once()

		result, err := doctorService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByIDDoctor(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		doctorRepository.On("GetByID", "1").Return(models.Doctor{}, nil).Once()

		result, err := doctorService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByID | Invalid", func(t *testing.T) {
		doctorRepository.On("GetByID", "0").Return(models.Doctor{}, errors.New("whoops")).Once()

		result, err := doctorService.GetByID("0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestCreateDoctor(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		doctorRepository.On("Create", models.DoctorRequest{}).Return(models.Doctor{}, nil).Once()

		result, err := doctorService.Create(models.DoctorRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create | Invalid", func(t *testing.T) {
		doctorRepository.On("Create", models.DoctorRequest{}).Return(models.Doctor{}, errors.New("oops")).Once()

		result, err := doctorService.Create(models.DoctorRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}


func TestUpdateDoctor(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		doctorRepository.On("Update", models.DoctorRequest{}, "1").Return(models.Doctor{}, nil).Once()

		result, err := doctorService.Update(models.DoctorRequest{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update | Invalid", func(t *testing.T) {
		doctorRepository.On("Update", models.DoctorRequest{}, "0").Return(models.Doctor{}, errors.New("oops")).Once()

		result, err := doctorService.Update(models.DoctorRequest{}, "0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDeleteDoctor(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		doctorRepository.On("Delete", "1").Return(nil).Once()

		err := doctorService.Delete("1")

		assert.Nil(t, err)
	})

	t.Run("Delete | Invalid", func(t *testing.T) {
		doctorRepository.On("Delete", "0").Return(errors.New("whoops")).Once()

		err := doctorService.Delete("0")

		assert.NotNil(t, err)
	})
}
