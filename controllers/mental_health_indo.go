package controllers

import (
	"MentalHealthCare/database"
	"MentalHealthCare/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetMentalHealthInfo(c *gin.Context) {
	var info []models.Information
	if err := database.DB.Find(&info).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, info)
}


func GetMentalHealthInfoByID(c *gin.Context) {
	var info models.Information
	infoID := c.Param("id")

	if err := database.DB.First(&info, infoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Information not found"})
		return
	}
	c.JSON(http.StatusOK, info)
}


func CreateMentalHealthInfo(c *gin.Context) {
	var info models.Information
	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&info).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, info)
}


func UpdateMentalHealthInfo(c *gin.Context) {
	var info models.Information
	infoID := c.Param("id")

	
	if err := database.DB.First(&info, infoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Information not found"})
		return
	}


	if err := c.ShouldBindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	if err := database.DB.Save(&info).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Information updated successfully", "info": info})
}


func DeleteMentalHealthInfo(c *gin.Context) {
	var info models.Information
	infoID := c.Param("id")

	
	if err := database.DB.First(&info, infoID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Information not found"})
		return
	}

	
	if err := database.DB.Delete(&info).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Information deleted successfully"})
}
