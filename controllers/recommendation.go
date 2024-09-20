package controllers

import (
	"MentalHealthCare/database"
	"MentalHealthCare/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetRecommendations retrieves all recommendations
func GetRecommendations(c echo.Context) error {
	var recommendations []models.Recommendation
	if err := database.DB.Find(&recommendations).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, recommendations)
}

// GetRecommendationByID retrieves a specific recommendation by ID
func GetRecommendationByID(c echo.Context) error {
	recommendationID := c.Param("id")
	var recommendation models.Recommendation

	if err := database.DB.First(&recommendation, recommendationID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Recommendation not found"})
	}
	return c.JSON(http.StatusOK, recommendation)
}

// CreateRecommendation creates a new recommendation
func CreateRecommendation(c echo.Context) error {
	var recommendation models.Recommendation
	if err := c.Bind(&recommendation); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := database.DB.Create(&recommendation).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, recommendation)
}

// UpdateRecommendation updates an existing recommendation by ID
func UpdateRecommendation(c echo.Context) error {
	recommendationID := c.Param("id")
	var recommendation models.Recommendation

	if err := database.DB.First(&recommendation, recommendationID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Recommendation not found"})
	}

	if err := c.Bind(&recommendation); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if err := database.DB.Save(&recommendation).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"message": "Recommendation updated successfully", "recommendation": recommendation})
}

// DeleteRecommendation deletes a recommendation by ID
func DeleteRecommendation(c echo.Context) error {
	recommendationID := c.Param("id")
	var recommendation models.Recommendation

	if err := database.DB.First(&recommendation, recommendationID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Recommendation not found"})
	}

	if err := database.DB.Delete(&recommendation).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "Recommendation deleted successfully"})
}
