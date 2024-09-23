// package repositories_test

// import (
// 	"MentalHealthCare/models"
// 	"errors"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// var recommendationRepository *mocks.RecommendationRepository // Assuming you have a mock setup
// var recommendationService *RecommendationService              // Assuming you have a Recommendation service

// func TestMain(m *testing.M) {
// 	recommendationRepository = new(mocks.RecommendationRepository)
// 	recommendationService = NewRecommendationService(recommendationRepository) // Initialize your service with the mock repository
// 	m.Run()
// }

// func TestGetAll(t *testing.T) {
// 	t.Run("GetAll | Valid", func(t *testing.T) {
// 		recommendationRepository.On("GetAll").Return([]models.Recommendation{}, nil).Once()

// 		result, err := recommendationService.GetAll()

// 		assert.NotNil(t, result)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("GetAll | Invalid", func(t *testing.T) {
// 		recommendationRepository.On("GetAll").Return([]models.Recommendation{}, errors.New("error")).Once()

// 		result, err := recommendationService.GetAll()

// 		assert.NotNil(t, result)
// 		assert.NotNil(t, err)
// 	})
// }

// func TestGetByID(t *testing.T) {
// 	t.Run("GetByID | Valid", func(t *testing.T) {
// 		recommendationRepository.On("GetByID", "1").Return(models.Recommendation{}, nil).Once()

// 		result, err := recommendationService.GetByID("1")

// 		assert.NotNil(t, result)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("GetByID | Invalid", func(t *testing.T) {
// 		recommendationRepository.On("GetByID", "0").Return(models.Recommendation{}, errors.New("whoops")).Once()

// 		result, err := recommendationService.GetByID("0")

// 		assert.NotNil(t, result)
// 		assert.NotNil(t, err)
// 	})
// }

// func TestCreate(t *testing.T) {
// 	t.Run("Create | Valid", func(t *testing.T) {
// 		recReq := models.Recommendation{Description: "Recommended treatment", DoctorID: 1, UserID: 2}
// 		recommendationRepository.On("Create", recReq).Return(recReq, nil).Once()

// 		result, err := recommendationService.Create(recReq)

// 		assert.NotNil(t, result)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Create | Invalid", func(t *testing.T) {
// 		recReq := models.Recommendation{Description: "Recommended treatment", DoctorID: 1, UserID: 2}
// 		recommendationRepository.On("Create", recReq).Return(models.Recommendation{}, errors.New("whoops")).Once()

// 		result, err := recommendationService.Create(recReq)

// 		assert.NotNil(t, result)
// 		assert.NotNil(t, err)
// 	})
// }

// func TestUpdate(t *testing.T) {
// 	t.Run("Update | Valid", func(t *testing.T) {
// 		recReq := models.Recommendation{Description: "Updated recommendation", DoctorID: 1, UserID: 2}
// 		recommendationRepository.On("Update", recReq, "1").Return(recReq, nil).Once()

// 		result, err := recommendationService.Update(recReq, "1")

// 		assert.NotNil(t, result)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Update | Invalid", func(t *testing.T) {
// 		recReq := models.Recommendation{Description: "Updated recommendation", DoctorID: 1, UserID: 2}
// 		recommendationRepository.On("Update", recReq, "0").Return(models.Recommendation{}, errors.New("whoops")).Once()

// 		result, err := recommendationService.Update(recReq, "0")

// 		assert.NotNil(t, result)
// 		assert.NotNil(t, err)
// 	})
// }

// func TestDelete(t *testing.T) {
// 	t.Run("Delete | Valid", func(t *testing.T) {
// 		recommendationRepository.On("Delete", "1").Return(nil).Once()

// 		err := recommendationService.Delete("1")

// 		assert.Nil(t, err)
// 	})

// 	t.Run("Delete | Invalid", func(t *testing.T) {
// 		recommendationRepository.On("Delete", "0").Return(errors.New("whoops")).Once()

// 		err := recommendationService.Delete("0")

// 		assert.NotNil(t, err)
// 	})
// }

///

package repositories_test

import (
	"MentalHealthCare/models"
	"MentalHealthCare/repositories/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecommendationRepository_Create(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewRecommendationRepository(t)

	// Define a sample recommendation to create
	recommendationReq := models.Recommendation{
		Description:      "This is a Description message.",
		DoctorID:  102,
		UserID : 1,
	}

	// Define the expected result
	expectedRecommendation := recommendationReq

	// Set up the expectation
	mockRepo.On("Create", recommendationReq).Return(expectedRecommendation, nil)

	// Call the method
	result, err := mockRepo.Create(recommendationReq)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedRecommendation, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestRecommendationRepository_Create_Error(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewRecommendationRepository(t)

	// Define a sample recommendation to create
	recommendationReq := models.Recommendation{
		Description:      "This is a Description message.",
		DoctorID:  102,
		UserID : 1,
	}

	// Define the error we expect
	expectedError := errors.New("failed to create recommendation")

	// Set up the expectation
	mockRepo.On("Create", recommendationReq).Return(models.Recommendation{}, expectedError)

	// Call the method
	result, err := mockRepo.Create(recommendationReq)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, models.Recommendation{}, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestRecommendationRepository_Delete(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewRecommendationRepository(t)

	// Define the ID to delete
	recommendationID := "1"

	// Set up the expectation
	mockRepo.On("Delete", recommendationID).Return(nil)

	// Call the method
	err := mockRepo.Delete(recommendationID)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestRecommendationRepository_GetAll(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewRecommendationRepository(t)

	// Define the expected recommendations
	expectedRecommendations := []models.Recommendation{
		{Description: "Description 1", DoctorID: 102, UserID: 1},
		{Description: "Description 2", DoctorID: 103, UserID: 2},
	}

	// Set up the expectation
	mockRepo.On("GetAll").Return(expectedRecommendations, nil)

	// Call the method
	result, err := mockRepo.GetAll()

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedRecommendations, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestRecommendationRepository_GetByID(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewRecommendationRepository(t)

	// Define the ID to fetch
	recommendationID := "1"

	// Define the expected recommendation
	expectedRecommendation := models.Recommendation{
		Description:      "This is a Description message.",
		DoctorID:  102,
		UserID : 1,
	}

	// Set up the expectation
	mockRepo.On("GetByID", recommendationID).Return(expectedRecommendation, nil)

	// Call the method
	result, err := mockRepo.GetByID(recommendationID)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedRecommendation, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestRecommendationRepository_Update(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewRecommendationRepository(t)

	// Define a sample recommendation to update
	recommendationReq := models.Recommendation{
		Description:      "This is a Description message.",
		DoctorID:  102,
		UserID : 1,
	}

	// Define the expected updated recommendation
	expectedRecommendation := recommendationReq

	// Set up the expectation
	mockRepo.On("Update", recommendationReq, "1").Return(expectedRecommendation, nil)

	// Call the method
	result, err := mockRepo.Update(recommendationReq, "1")

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedRecommendation, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}
