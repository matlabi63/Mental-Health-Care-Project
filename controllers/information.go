package controllers

import (
	"MentalHealthCare/database"
	"MentalHealthCare/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetMentalHealthInfo retrieves all mental health information
func GetMentalHealthInfo(c echo.Context) error {
	var info []models.Information
	if err := database.DB.Find(&info).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, info)
}

// GetMentalHealthInfoByID retrieves specific mental health information by ID
func GetMentalHealthInfoByID(c echo.Context) error {
	infoID := c.Param("id")
	var info models.Information

	if err := database.DB.First(&info, infoID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Information not found"})
	}
	return c.JSON(http.StatusOK, info)
}

// CreateMentalHealthInfo creates new mental health information
func CreateMentalHealthInfo(c echo.Context) error {
	var info models.Information
	if err := c.Bind(&info); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := database.DB.Create(&info).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, info)
}

// UpdateMentalHealthInfo updates existing mental health information by ID
func UpdateMentalHealthInfo(c echo.Context) error {
	infoID := c.Param("id")
	var info models.Information

	if err := database.DB.First(&info, infoID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Information not found"})
	}

	if err := c.Bind(&info); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := database.DB.Save(&info).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Information updated successfully", "info": info})
}

// DeleteMentalHealthInfo deletes mental health information by ID
func DeleteMentalHealthInfo(c echo.Context) error {
	infoID := c.Param("id")
	var info models.Information

	if err := database.DB.First(&info, infoID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Information not found"})
	}

	if err := database.DB.Delete(&info).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Information deleted successfully"})
}
