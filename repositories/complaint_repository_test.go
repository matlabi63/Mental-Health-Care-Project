package repositories_test

import (
	"MentalHealthCare/models"
	"MentalHealthCare/repositories/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComplaintRepository_Create(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewComplaintRepository(t)

	// Define a sample complaint to create
	complaintReq := models.Complaint{
		UserID:  101,
		Details: "This is a complaint message.",
	}

	// Define the expected result
	expectedComplaint := complaintReq

	// Set up the expectation
	mockRepo.On("Create", complaintReq).Return(expectedComplaint, nil)

	// Call the method
	result, err := mockRepo.Create(complaintReq)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedComplaint, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestComplaintRepository_Create_Error(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewComplaintRepository(t)

	// Define a sample complaint to create
	complaintReq := models.Complaint{
		UserID:  101,
		Details: "This is a complaint message.",
	}

	// Define the error we expect
	expectedError := errors.New("failed to create complaint")

	// Set up the expectation
	mockRepo.On("Create", complaintReq).Return(models.Complaint{}, expectedError)

	// Call the method
	result, err := mockRepo.Create(complaintReq)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, models.Complaint{}, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestComplaintRepository_Delete(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewComplaintRepository(t)

	// Define the ID to delete
	complaintID := "1"

	// Set up the expectation
	mockRepo.On("Delete", complaintID).Return(nil)

	// Call the method
	err := mockRepo.Delete(complaintID)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestComplaintRepository_GetAll(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewComplaintRepository(t)

	// Define the expected complaints
	expectedComplaints := []models.Complaint{
		{ID: 1, UserID: 101, Details: "Complaint 1"},
		{ID: 2, UserID: 102, Details: "Complaint 2"},
	}

	// Set up the expectation
	mockRepo.On("GetAll").Return(expectedComplaints, nil)

	// Call the method
	result, err := mockRepo.GetAll()

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedComplaints, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestComplaintRepository_GetByID(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewComplaintRepository(t)

	// Define the ID to fetch
	complaintID := "1"

	// Define the expected complaint
	expectedComplaint := models.Complaint{
		UserID:  101,
		Details: "This is a complaint message.",
	}

	// Set up the expectation
	mockRepo.On("GetByID", complaintID).Return(expectedComplaint, nil)

	// Call the method
	result, err := mockRepo.GetByID(complaintID)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedComplaint, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestComplaintRepository_Update(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewComplaintRepository(t)

	// Define a sample complaint to update
	complaintReq := models.Complaint{
		UserID:  101,
		Details: "This is a complaint message.",
	}

	// Define the expected updated complaint
	expectedComplaint := complaintReq

	// Set up the expectation
	mockRepo.On("Update", complaintReq, "1").Return(expectedComplaint, nil)

	// Call the method
	result, err := mockRepo.Update(complaintReq, "1")

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedComplaint, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}
