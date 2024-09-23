// package repositories_test

// import (
// 	"MentalHealthCare/models"
// 	"errors"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// var doctorRepository *mocks.DoctorRepository // Assuming you have a mock setup
// var doctorService *DoctorService              // Assuming you have a Doctor service

// func TestMain(m *testing.M) {
// 	doctorRepository = new(mocks.DoctorRepository)
// 	doctorService = NewDoctorService(doctorRepository) // Initialize your service with the mock repository
// 	m.Run()
// }

// func TestGetAll(t *testing.T) {
// 	t.Run("GetAll | Valid", func(t *testing.T) {
// 		doctorRepository.On("GetAll").Return([]models.Doctor{}, nil).Once()

// 		result, err := doctorService.GetAll()

// 		assert.NotNil(t, result)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("GetAll | Invalid", func(t *testing.T) {
// 		doctorRepository.On("GetAll").Return([]models.Doctor{}, errors.New("error")).Once()

// 		result, err := doctorService.GetAll()

// 		assert.NotNil(t, result)
// 		assert.NotNil(t, err)
// 	})
// }

// func TestGetByID(t *testing.T) {
// 	t.Run("GetByID | Valid", func(t *testing.T) {
// 		doctorRepository.On("GetByID", "1").Return(models.Doctor{}, nil).Once()

// 		result, err := doctorService.GetByID("1")

// 		assert.NotNil(t, result)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("GetByID | Invalid", func(t *testing.T) {
// 		doctorRepository.On("GetByID", "0").Return(models.Doctor{}, errors.New("whoops")).Once()

// 		result, err := doctorService.GetByID("0")

// 		assert.NotNil(t, result)
// 		assert.NotNil(t, err)
// 	})
// }

// func TestCreate(t *testing.T) {
// 	t.Run("Create | Valid", func(t *testing.T) {
// 		doctorReq := models.Doctor{Name: "Dr. Smith", Specialty: "Cardiology"}
// 		doctorRepository.On("Create", doctorReq).Return(doctorReq, nil).Once()

// 		result, err := doctorService.Create(doctorReq)

// 		assert.NotNil(t, result)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Create | Invalid", func(t *testing.T) {
// 		doctorReq := models.Doctor{Name: "Dr. Smith", Specialty: "Cardiology"}
// 		doctorRepository.On("Create", doctorReq).Return(models.Doctor{}, errors.New("whoops")).Once()

// 		result, err := doctorService.Create(doctorReq)

// 		assert.NotNil(t, result)
// 		assert.NotNil(t, err)
// 	})
// }

// func TestUpdate(t *testing.T) {
// 	t.Run("Update | Valid", func(t *testing.T) {
// 		doctorReq := models.Doctor{Name: "Dr. Smith", Specialty: "Cardiology"}
// 		doctorRepository.On("Update", doctorReq, "1").Return(doctorReq, nil).Once()

// 		result, err := doctorService.Update(doctorReq, "1")

// 		assert.NotNil(t, result)
// 		assert.Nil(t, err)
// 	})

// 	t.Run("Update | Invalid", func(t *testing.T) {
// 		doctorReq := models.Doctor{Name: "Dr. Smith", Specialty: "Cardiology"}
// 		doctorRepository.On("Update", doctorReq, "0").Return(models.Doctor{}, errors.New("whoops")).Once()

// 		result, err := doctorService.Update(doctorReq, "0")

// 		assert.NotNil(t, result)
// 		assert.NotNil(t, err)
// 	})
// }

// func TestDelete(t *testing.T) {
// 	t.Run("Delete | Valid", func(t *testing.T) {
// 		doctorRepository.On("Delete", "1").Return(nil).Once()

// 		err := doctorService.Delete("1")

// 		assert.Nil(t, err)
// 	})

// 	t.Run("Delete | Invalid", func(t *testing.T) {
// 		doctorRepository.On("Delete", "0").Return(errors.New("whoops")).Once()

// 		err := doctorService.Delete("0")

// 		assert.NotNil(t, err)
// 	})
// }

/////

package repositories_test

import (
	"MentalHealthCare/models"
	"MentalHealthCare/repositories/mocks"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoctorRepository_Create(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewDoctorRepository(t)

	// Define a sample doctor to create
	doctorReq := models.Doctor{
		UserID:    101,
		Name:  "Dr. Smith",
		Specialty: "Surgion",
	}

	// Define the expected result
	expectedDoctor := doctorReq

	// Set up the expectation
	mockRepo.On("Create", doctorReq).Return(expectedDoctor, nil)

	// Call the method
	result, err := mockRepo.Create(doctorReq)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedDoctor, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestDoctorRepository_Create_Error(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewDoctorRepository(t)

	// Define a sample doctor to create
	doctorReq := models.Doctor{
		UserID:    101,
		Name:  "Dr. Smith",
		Specialty: "Surgion",
	}

	// Define the error we expect
	expectedError := errors.New("failed to create doctor")

	// Set up the expectation
	mockRepo.On("Create", doctorReq).Return(models.Doctor{}, expectedError)

	// Call the method
	result, err := mockRepo.Create(doctorReq)

	// Assert the result
	assert.Error(t, err)
	assert.Equal(t, expectedError, err)
	assert.Equal(t, models.Doctor{}, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestDoctorRepository_Delete(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewDoctorRepository(t)

	// Define the ID to delete
	doctorID := "1"

	// Set up the expectation
	mockRepo.On("Delete", doctorID).Return(nil)

	// Call the method
	err := mockRepo.Delete(doctorID)

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestDoctorRepository_GetAll(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewDoctorRepository(t)

	// Define the expected doctors
	expectedDoctors := []models.Doctor{
		{ID: 1, Name: "Dr. Smith", Specialty: "Surgion"},
		{ID: 2, Name: "Dr. Johnson", Specialty: "Dentist"},
	}

	// Set up the expectation
	mockRepo.On("GetAll").Return(expectedDoctors, nil)

	// Call the method
	result, err := mockRepo.GetAll()

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedDoctors, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestDoctorRepository_GetByID(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewDoctorRepository(t)

	// Define the ID to fetch
	doctorID := "1"

	// Define the expected doctor
	expectedDoctor := models.Doctor{
		UserID:    101,
		Name:  "Dr. Smith",
		Specialty: "Surgion",
	}

	// Set up the expectation
	mockRepo.On("GetByID", doctorID).Return(expectedDoctor, nil)

	// Call the method
	result, err := mockRepo.GetByID(doctorID)

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedDoctor, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}

func TestDoctorRepository_Update(t *testing.T) {
	// Initialize the mock repository
	mockRepo := mocks.NewDoctorRepository(t)

	// Define a sample doctor to update
	doctorReq := models.Doctor{
		UserID:    101,
		Name:  "Dr. Smith",
		Specialty: "Surgion",
	}

	// Define the expected updated doctor
	expectedDoctor := doctorReq

	// Set up the expectation
	mockRepo.On("Update", doctorReq, "1").Return(expectedDoctor, nil)

	// Call the method
	result, err := mockRepo.Update(doctorReq, "1")

	// Assert the result
	assert.NoError(t, err)
	assert.Equal(t, expectedDoctor, result)

	// Assert that the expectations were met
	mockRepo.AssertExpectations(t)
}
