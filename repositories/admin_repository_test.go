package repositories_test

import (
	"MentalHealthCare/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllAdmin(t *testing.T) {
	t.Run("GetAll | Valid", func(t *testing.T) {
		adminRepository.On("GetAll").Return([]models.Admin{}, nil).Once()

		result, err := adminService.GetAll()

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetAll | Invalid", func(t *testing.T) {
		adminRepository.On("GetAll").Return([]models.Admin{}, errors.New("error")).Once()

		result, err := adminService.GetAll()

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestGetByIDAdmin(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		adminRepository.On("GetByID", "1").Return(models.Admin{}, nil).Once()

		result, err := adminService.GetByID("1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("GetByID | Invalid", func(t *testing.T) {
		adminRepository.On("GetByID", "0").Return(models.Admin{}, errors.New("whoops")).Once()

		result, err := adminService.GetByID("0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}


func TestCreateAdmin(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		adminRepository.On("Create", models.AdminRequest{}).Return(models.Admin{}, nil).Once()

		result, err := adminService.Create(models.AdminRequest{})

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Create | Invalid", func(t *testing.T) {
		adminRepository.On("Create", models.AdminRequest{}).Return(models.Admin{}, errors.New("whoops")).Once()

		result, err := adminService.Create(models.AdminRequest{})

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestUpdateAdmin(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		adminRepository.On("Update", models.AdminRequest{}, "1").Return(models.Admin{}, nil).Once()

		result, err := adminService.Update(models.AdminRequest{}, "1")

		assert.NotNil(t, result)
		assert.Nil(t, err)
	})

	t.Run("Update | Invalid", func(t *testing.T) {
		adminRepository.On("Update", models.AdminRequest{}, "0").Return(models.Admin{}, errors.New("whoops")).Once()

		result, err := adminService.Update(models.AdminRequest{}, "0")

		assert.NotNil(t, result)
		assert.NotNil(t, err)
	})
}

func TestDeleteAdmin(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		adminRepository.On("Delete", "1").Return(nil).Once()

		err := adminService.Delete("1")

		assert.Nil(t, err)
	})

	t.Run("Delete | Invalid", func(t *testing.T) {
		adminRepository.On("Delete", "0").Return(errors.New("whoops")).Once()

		err := adminService.Delete("0")

		assert.NotNil(t, err)
	})
}
