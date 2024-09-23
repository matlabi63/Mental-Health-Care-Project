package repositories_test

import (
	"MentalHealthCare/models"
	"MentalHealthCare/repositories/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdminRepository_Create(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewAdminRepository(t)

	// Define a sample admin to create
	adminReq := models.Admin{
		UserID:      1,
    ManageUsers: true,
	}

	// Define the expected result
	expectedAdmin := adminReq

	// Set up the expectation
	mockRepo.On("Create", adminReq).Return(expectedAdmin, nil)

	// Call the method
	result, err := mockRepo.Create(adminReq)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedAdmin, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestAdminRepository_Create_Error(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewAdminRepository(t)

	// Define a sample admin to create
	adminReq := models.Admin{
		UserID:      1,
    ManageUsers: true,
	}

	// Define the error we expect
	expectedError := errors.New("failed to create admin")

	// Set up the expectation
	mockRepo.On("Create", adminReq).Return(models.Admin{}, expectedError)

	// Call the method
	result, err := mockRepo.Create(adminReq)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, models.Admin{}, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestAdminRepository_Delete(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewAdminRepository(t)

	// Define the ID to delete
	adminID := "1"

	// Set up the expectation
	mockRepo.On("Delete", adminID).Return(nil)

	// Call the method
	err := mockRepo.Delete(adminID)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestAdminRepository_GetAll(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewAdminRepository(t)

	// Define the expected admins
	expectedAdmins := []models.Admin{
		{ID: 1, UserID : 1, ManageUsers: true},
    {ID: 2, UserID : 2, ManageUsers: false},
	}

	// Set up the expectation
	mockRepo.On("GetAll").Return(expectedAdmins, nil)

	// Call the method
	result, err := mockRepo.GetAll()

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedAdmins, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestAdminRepository_GetByID(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewAdminRepository(t)

	// Define the ID to fetch
	adminID := "1"

	// Define the expected admin
	expectedAdmin := models.Admin{
		UserID:      1,
    ManageUsers: true,
	}

	// Set up the expectation
	mockRepo.On("GetByID", adminID).Return(expectedAdmin, nil)

	// Call the method
	result, err := mockRepo.GetByID(adminID)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedAdmin, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestAdminRepository_Update(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewAdminRepository(t)

	// Define a sample admin to update
	adminReq := models.Admin{
		UserID:      1,
    ManageUsers: true,
	}

	// Define the expected updated admin
	expectedAdmin := adminReq

	// Set up the expectation
	mockRepo.On("Update", adminReq, "1").Return(expectedAdmin, nil)

	// Call the method
	result, err := mockRepo.Update(adminReq, "1")

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedAdmin, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}
