package controllers

import (
	"MentalHealthCare/database"
	"MentalHealthCare/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetDoctors retrieves all doctors from the database
func GetDoctors(c echo.Context) error {
	var doctors []models.Doctor
	if err := database.DB.Find(&doctors).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, doctors)
}

// GetDoctorByID retrieves a doctor by its ID
func GetDoctorByID(c echo.Context) error {
	doctorID := c.Param("id")
	var doctor models.Doctor

	if err := database.DB.First(&doctor, doctorID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Doctor not found"})
	}
	return c.JSON(http.StatusOK, doctor)
}

// CreateDoctor creates a new doctor
func CreateDoctor(c echo.Context) error {
	var doctor models.Doctor
	if err := c.Bind(&doctor); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Check if the associated user exists
	var user models.User
	if err := database.DB.First(&user, doctor.UserID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "User not found"})
	}

	// Create doctor entry
	if err := database.DB.Create(&doctor).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, doctor)
}

// UpdateDoctor updates an existing doctor by ID
func UpdateDoctor(c echo.Context) error {
	doctorID := c.Param("id")
	var doctor models.Doctor

	// Find the doctor by ID
	if err := database.DB.First(&doctor, doctorID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Doctor not found"})
	}

	// Bind new data
	if err := c.Bind(&doctor); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// Check if the user exists if UserID is provided
	if doctor.UserID != 0 {
		var user models.User
		if err := database.DB.First(&user, doctor.UserID).Error; err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "User not found"})
		}
	}

	// Save the updated doctor record
	if err := database.DB.Save(&doctor).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Doctor updated successfully", "doctor": doctor})
}

// DeleteDoctor deletes a doctor by ID
func DeleteDoctor(c echo.Context) error {
	doctorID := c.Param("id")
	var doctor models.Doctor

	// Find the doctor by ID
	if err := database.DB.First(&doctor, doctorID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Doctor not found"})
	}

	// Delete the doctor record
	if err := database.DB.Delete(&doctor).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Doctor deleted successfully"})
}
