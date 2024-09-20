package controllers

import (
	"MentalHealthCare/database"
	"MentalHealthCare/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetComplaints retrieves all complaints from the database
func GetComplaints(c echo.Context) error {
	var complaints []models.Complaint
	if err := database.DB.Find(&complaints).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, complaints)
}

// GetComplaintByID retrieves a complaint by its ID
func GetComplaintByID(c echo.Context) error {
	complaintID := c.Param("id")
	var complaint models.Complaint

	if err := database.DB.First(&complaint, complaintID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Complaint not found"})
	}
	return c.JSON(http.StatusOK, complaint)
}

// CreateComplaint creates a new complaint
func CreateComplaint(c echo.Context) error {
	var complaint models.Complaint
	if err := c.Bind(&complaint); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := database.DB.Create(&complaint).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, complaint)
}

// UpdateComplaint updates an existing complaint
func UpdateComplaint(c echo.Context) error {
	complaintID := c.Param("id")
	var complaint models.Complaint

	if err := database.DB.First(&complaint, complaintID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Complaint not found"})
	}

	if err := c.Bind(&complaint); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := database.DB.Save(&complaint).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Complaint updated successfully", "complaint": complaint})
}

// DeleteComplaint deletes a complaint by its ID
func DeleteComplaint(c echo.Context) error {
	complaintID := c.Param("id")
	var complaint models.Complaint

	if err := database.DB.First(&complaint, complaintID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Complaint not found"})
	}

	if err := database.DB.Delete(&complaint).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Complaint deleted successfully"})
}
