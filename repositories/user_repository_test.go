// package repositories_test

// import (
// 	"MentalHealthCare/models"
// 	"MentalHealthCare/repositories/mocks"
// 	"errors"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestUserRepository_Create(t *testing.T) {
// 	// Initialize the mock repository
// 	mockRepo := mocks.NewUserRepository(t)

// 	// Define a sample user to create
// 	userReq := models.User{
// 		ID:   "1",
// 		Name: "Test User",
// 		Email: "testuser@example.com",
// 	}

// 	// Define the expected result
// 	expectedUser := userReq

// 	// Set up the expectation
// 	mockRepo.On("Create", userReq).Return(expectedUser, nil)

// 	// Call the method
// 	result, err := mockRepo.Create(userReq)

// 	// Assert the result
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedUser, result)

// 	// Assert that the expectations were met
// 	mockRepo.AssertExpectations(t)
// }

// func TestUserRepository_Create_Error(t *testing.T) {
// 	// Initialize the mock repository
// 	mockRepo := mocks.NewUserRepository(t)

// 	// Define a sample user to create
// 	userReq := models.User{
// 		ID:   "1",
// 		Name: "Test User",
// 		Email: "testuser@example.com",
// 	}

// 	// Define the error we expect
// 	expectedError := errors.New("failed to create user")

// 	// Set up the expectation
// 	mockRepo.On("Create", userReq).Return(models.User{}, expectedError)

// 	// Call the method
// 	result, err := mockRepo.Create(userReq)

// 	// Assert the result
// 	assert.Error(t, err)
// 	assert.Equal(t, expectedError, err)
// 	assert.Equal(t, models.User{}, result)

// 	// Assert that the expectations were met
// 	mockRepo.AssertExpectations(t)
// }

// func TestUserRepository_Delete(t *testing.T) {
// 	// Initialize the mock repository
// 	mockRepo := mocks.NewUserRepository(t)

// 	// Define the ID to delete
// 	userID := "1"

// 	// Set up the expectation
// 	mockRepo.On("Delete", userID).Return(nil)

// 	// Call the method
// 	err := mockRepo.Delete(userID)

// 	// Assert the result
// 	assert.NoError(t, err)

// 	// Assert that the expectations were met
// 	mockRepo.AssertExpectations(t)
// }

// func TestUserRepository_GetAll(t *testing.T) {
// 	// Initialize the mock repository
// 	mockRepo := mocks.NewUserRepository(t)

// 	// Define the expected users
// 	expectedUsers := []models.User{
// 		{ID: "1", Name: "Test User 1", Email: "test1@example.com"},
// 		{ID: "2", Name: "Test User 2", Email: "test2@example.com"},
// 	}

// 	// Set up the expectation
// 	mockRepo.On("GetAll").Return(expectedUsers, nil)

// 	// Call the method
// 	result, err := mockRepo.GetAll()

// 	// Assert the result
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedUsers, result)

// 	// Assert that the expectations were met
// 	mockRepo.AssertExpectations(t)
// }

// func TestUserRepository_GetByID(t *testing.T) {
// 	// Initialize the mock repository
// 	mockRepo := mocks.NewUserRepository(t)

// 	// Define the ID to fetch
// 	userID := "1"

// 	// Define the expected user
// 	expectedUser := models.User{
// 		ID:   "1",
// 		Name: "Test User",
// 		Email: "testuser@example.com",
// 	}

// 	// Set up the expectation
// 	mockRepo.On("GetByID", userID).Return(expectedUser, nil)

// 	// Call the method
// 	result, err := mockRepo.GetByID(userID)

// 	// Assert the result
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedUser, result)

// 	// Assert that the expectations were met
// 	mockRepo.AssertExpectations(t)
// }

// func TestUserRepository_Update(t *testing.T) {
// 	// Initialize the mock repository
// 	mockRepo := mocks.NewUserRepository(t)

// 	// Define a sample user to update
// 	userReq := models.User{
// 		ID:    "1",
// 		Name:  "Updated User",
// 		Email: "updateduser@example.com",
// 	}

// 	// Define the expected updated user
// 	expectedUser := userReq

// 	// Set up the expectation
// 	mockRepo.On("Update", userReq, "1").Return(expectedUser, nil)

// 	// Call the method
// 	result, err := mockRepo.Update(userReq, "1")

// 	// Assert the result
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedUser, result)

// 	// Assert that the expectations were met
// 	mockRepo.AssertExpectations(t)
// }

////

package repositories_test

import (
	"MentalHealthCare/models"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is the mock for UserRepository interface
type MockUserRepository struct {
	mock.Mock
}

// Mock implementations for UserRepository methods
func (m *MockUserRepository) GetAll() ([]models.User, error) {
	args := m.Called()
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserRepository) GetByID(id string) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) Create(user models.User) (models.User, error) {
	args := m.Called(user)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) Update(user models.User, id string) (models.User, error) {
	args := m.Called(user, id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) Delete(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestGetAll(t *testing.T) {
	userRepo := new(MockUserRepository)
	userRepo.On("GetAll").Return([]models.User{
		{ID: 1, Name: "John Doe", Email: "john@example.com"},
	}, nil).Once()

	// Act
	result, err := userRepo.GetAll()

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result))
	userRepo.AssertExpectations(t)
}

func TestGetByID(t *testing.T) {
	userRepo := new(MockUserRepository)

	t.Run("GetByID | Valid", func(t *testing.T) {
		userRepo.On("GetByID", "1").Return(models.User{
			ID:    1,
			Name:  "John Doe",
			Email: "john@example.com",
		}, nil).Once()

		// Act
		result, err := userRepo.GetByID("1")

		// Assert
		assert.Nil(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "John Doe", result.Name)
		userRepo.AssertExpectations(t)
	})

	t.Run("GetByID | Invalid", func(t *testing.T) {
		userRepo.On("GetByID", "999").Return(models.User{}, errors.New("user not found")).Once()

		// Act
		result, err := userRepo.GetByID("999")

		// Assert
		assert.NotNil(t, err)
		assert.Equal(t, "user not found", err.Error())
		assert.Equal(t, models.User{}, result)
		userRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	userRepo := new(MockUserRepository)

	t.Run("Create | Valid", func(t *testing.T) {
		newUser := models.User{
			Name:  "Jane Doe",
			Email: "jane@example.com",
		}

		userRepo.On("Create", newUser).Return(newUser, nil).Once()

		// Act
		result, err := userRepo.Create(newUser)

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, "Jane Doe", result.Name)
		assert.Equal(t, "jane@example.com", result.Email)
		userRepo.AssertExpectations(t)
	})

	t.Run("Create | Invalid", func(t *testing.T) {
		newUser := models.User{
			Name:  "Jane Doe",
			Email: "jane@example.com",
		}

		userRepo.On("Create", newUser).Return(models.User{}, errors.New("creation failed")).Once()

		// Act
		result, err := userRepo.Create(newUser)

		// Assert
		assert.NotNil(t, err)
		assert.Equal(t, "creation failed", err.Error())
		assert.Equal(t, models.User{}, result)
		userRepo.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	userRepo := new(MockUserRepository)

	t.Run("Update | Valid", func(t *testing.T) {
		updatedUser := models.User{
			Name:  "John Updated",
			Email: "john_updated@example.com",
		}

		userRepo.On("Update", updatedUser, "1").Return(updatedUser, nil).Once()

		// Act
		result, err := userRepo.Update(updatedUser, "1")

		// Assert
		assert.Nil(t, err)
		assert.Equal(t, "John Updated", result.Name)
		assert.Equal(t, "john_updated@example.com", result.Email)
		userRepo.AssertExpectations(t)
	})

	t.Run("Update | Invalid", func(t *testing.T) {
		updatedUser := models.User{
			Name:  "John Updated",
			Email: "john_updated@example.com",
		}

		userRepo.On("Update", updatedUser, "999").Return(models.User{}, errors.New("update failed")).Once()

		// Act
		result, err := userRepo.Update(updatedUser, "999")

		// Assert
		assert.NotNil(t, err)
		assert.Equal(t, "update failed", err.Error())
		assert.Equal(t, models.User{}, result)
		userRepo.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	userRepo := new(MockUserRepository)

	t.Run("Delete | Valid", func(t *testing.T) {
		userRepo.On("Delete", "1").Return(nil).Once()

		// Act
		err := userRepo.Delete("1")

		// Assert
		assert.Nil(t, err)
		userRepo.AssertExpectations(t)
	})

	t.Run("Delete | Invalid", func(t *testing.T) {
		userRepo.On("Delete", "999").Return(errors.New("delete failed")).Once()

		// Act
		err := userRepo.Delete("999")

		// Assert
		assert.NotNil(t, err)
		assert.Equal(t, "delete failed", err.Error())
		userRepo.AssertExpectations(t)
	})
}
