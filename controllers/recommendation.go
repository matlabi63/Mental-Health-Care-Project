package controllers

import (
	"MentalHealthCare/database"
	"MentalHealthCare/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//get all recommendations
func GetRecommendations(c *gin.Context) {
	var recommendations []models.Recommendation
	if err := database.DB.Find(&recommendations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recommendations)
}

//get a specific recommendation by ID
func GetRecommendationByID(c *gin.Context) {
	recommendationID := c.Param("id")
	var recommendation models.Recommendation

	if err := database.DB.First(&recommendation, recommendationID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recommendation not found"})
		return
	}
	c.JSON(http.StatusOK, recommendation)
}

//creates a new recommendation
func CreateRecommendation(c *gin.Context) {
	var recommendation models.Recommendation
	if err := c.ShouldBindJSON(&recommendation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&recommendation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, recommendation)
}

// updates an existing recommendation by ID
func UpdateRecommendation(c *gin.Context) {
	recommendationID := c.Param("id")
	var recommendation models.Recommendation

	// Find the recommendation by ID
	if err := database.DB.First(&recommendation, recommendationID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recommendation not found"})
		return
	}

	// Bind the incoming JSON to the recommendation object
	if err := c.ShouldBindJSON(&recommendation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the updated recommendation details
	if err := database.DB.Save(&recommendation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recommendation updated successfully", "recommendation": recommendation})
}

// DeleteRecommendation deletes a recommendation by ID
func DeleteRecommendation(c *gin.Context) {
	recommendationID := c.Param("id")
	var recommendation models.Recommendation

	// Find the recommendation by ID
	if err := database.DB.First(&recommendation, recommendationID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Recommendation not found"})
		return
	}

	// Delete the recommendation
	if err := database.DB.Delete(&recommendation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Recommendation deleted successfully"})
}
