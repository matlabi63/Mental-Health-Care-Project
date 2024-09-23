// package repositories_test

// import (
// 	"MentalHealthCare/models"
// 	"errors"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// var informationRepository *mocks.InformationRepository // Assuming you have a mock setup
// var informationService *InformationService              // Assuming you have an Information service

// func TestMain(m *testing.M) {
// 	informationRepository = new(mocks.InformationRepository)
// 	informationService = NewInformationService(informationRepository) // Initialize your service with the mock repository
// 	m.Run()
// }

// func TestGetAll(t *testing.T) {
// 	t.Run("GetAll | Valid", func(t *testing.T) {
// 		informationRepository.On("GetAll").Return([]models.Information{}, nil).Once()

// 		result, err := informationService.GetAll()

// 		assert.NotNil(t, result)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("GetAll | Invalid", func(t *testing.T) {
// 		informationRepository.On("GetAll").Return([]models.Information{}, errors.New("error")).Once()

// 		result, err := informationService.GetAll()

// 		assert.NotNil(t, result)
// 		assert.NotNil(t, err)
// 	})
// }

// func TestGetByID(t *testing.T) {
// 	t.Run("GetByID | Valid", func(t *testing.T) {
// 		informationRepository.On("GetByID", "1").Return(models.Information{}, nil).Once()

// 		result, err := informationService.GetByID("1")

// 		assert.NotNil(t, result)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("GetByID | Invalid", func(t *testing.T) {
// 		informationRepository.On("GetByID", "0").Return(models.Information{}, errors.New("whoops")).Once()

// 		result, err := informationService.GetByID("0")

// 		assert.NotNil(t, result)
// 		assert.NotNil(t, err)
// 	})
// }

// func TestCreate(t *testing.T) {
// 	t.Run("Create | Valid", func(t *testing.T) {
// 		infoReq := models.Information{Title: "New Info", Content: "Content goes here", Author: "Author Name"}
// 		informationRepository.On("Create", infoReq).Return(infoReq, nil).Once()

// 		result, err := informationService.Create(infoReq)

// 		assert.NotNil(t, result)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Create | Invalid", func(t *testing.T) {
// 		infoReq := models.Information{Title: "New Info", Content: "Content goes here", Author: "Author Name"}
// 		informationRepository.On("Create", infoReq).Return(models.Information{}, errors.New("whoops")).Once()

// 		result, err := informationService.Create(infoReq)

// 		assert.NotNil(t, result)
// 		assert.NotNil(t, err)
// 	})
// }

// func TestUpdate(t *testing.T) {
// 	t.Run("Update | Valid", func(t *testing.T) {
// 		infoReq := models.Information{Title: "Updated Info", Content: "Updated content", Author: "Author Name"}
// 		informationRepository.On("Update", infoReq, "1").Return(infoReq, nil).Once()

// 		result, err := informationService.Update(infoReq, "1")

// 		assert.NotNil(t, result)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Update | Invalid", func(t *testing.T) {
// 		infoReq := models.Information{Title: "Updated Info", Content: "Updated content", Author: "Author Name"}
// 		informationRepository.On("Update", infoReq, "0").Return(models.Information{}, errors.New("whoops")).Once()

// 		result, err := informationService.Update(infoReq, "0")

// 		assert.NotNil(t, result)
// 		assert.NotNil(t, err)
// 	})
// }

// func TestDelete(t *testing.T) {
// 	t.Run("Delete | Valid", func(t *testing.T) {
// 		informationRepository.On("Delete", "1").Return(nil).Once()

// 		err := informationService.Delete("1")

// 		assert.Nil(t, err)
// 	})

// 	t.Run("Delete | Invalid", func(t *testing.T) {
// 		informationRepository.On("Delete", "0").Return(errors.New("whoops")).Once()

// 		err := informationService.Delete("0")

// 		assert.NotNil(t, err)
// 	})
// }

////

package repositories_test

import (
	"MentalHealthCare/models"
	"MentalHealthCare/repositories/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInformationRepository_Create(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewInformationRepository(t)

	// Define a sample information to create
	informationReq := models.Information{
		Title:   "Mental Health Awareness",
		Content: "Content about mental health awareness.",
		Author:	"Ahmad",
	}

	// Define the expected result
	expectedInformation := informationReq

	// Set up the expectation
	mockRepo.On("Create", informationReq).Return(expectedInformation, nil)

	// Call the method
	result, err := mockRepo.Create(informationReq)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedInformation, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestInformationRepository_Create_Error(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewInformationRepository(t)

	// Define a sample information to create
	informationReq := models.Information{
		Title:   "Mental Health Awareness",
		Content: "Content about mental health awareness.",
		Author:	"Ahmad",
	}

	// Define the error we expect
	expectedError := errors.New("failed to create information")

	// Set up the expectation
	mockRepo.On("Create", informationReq).Return(models.Information{}, expectedError)

	// Call the method
	result, err := mockRepo.Create(informationReq)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, models.Information{}, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestInformationRepository_Delete(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewInformationRepository(t)

	// Define the ID to delete
	informationID := "1"

	// Set up the expectation
	mockRepo.On("Delete", informationID).Return(nil)

	// Call the method
	err := mockRepo.Delete(informationID)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestInformationRepository_GetAll(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewInformationRepository(t)

	// Define the expected information
	expectedInformation := []models.Information{
		{Title: "Mental Health Awareness", Content: "Content about mental health awareness.", Author:	"Ahmad",},
		{Title: "Stress Management", Content: "Content about stress management.", Author:	"Mahmood",},
	}

	// Set up the expectation
	mockRepo.On("GetAll").Return(expectedInformation, nil)

	// Call the method
	result, err := mockRepo.GetAll()

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedInformation, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestInformationRepository_GetByID(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewInformationRepository(t)

	// Define the ID to fetch
	informationID := "1"

	// Define the expected information
	expectedInformation := models.Information{
		Title:   "Mental Health Awareness",
		Content: "Content about mental health awareness.",
		Author:	"Ahmad",
	}

	// Set up the expectation
	mockRepo.On("GetByID", informationID).Return(expectedInformation, nil)

	// Call the method
	result, err := mockRepo.GetByID(informationID)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedInformation, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestInformationRepository_Update(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewInformationRepository(t)

	// Define a sample information to update
	informationReq := models.Information{
		Title:   "Updated Mental Health Awareness",
		Content: "Updated content about mental health awareness.",
		Author:	"Ahmad",
	}

	// Define the expected updated information
	expectedInformation := informationReq

	// Set up the expectation
	mockRepo.On("Update", informationReq, "1").Return(expectedInformation, nil)

	// Call the method
	result, err := mockRepo.Update(informationReq, "1")

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedInformation, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}
