package controllers

import (
	"MentalHealthCare/database"
	"MentalHealthCare/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetComplaints(c *gin.Context) {
	var complaints []models.Complaint
	if err := database.DB.Find(&complaints).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, complaints)
}


func GetComplaintByID(c *gin.Context) {
	complaintID := c.Param("id")
	var complaint models.Complaint

	if err := database.DB.First(&complaint, complaintID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Complaint not found"})
		return
	}
	c.JSON(http.StatusOK, complaint)
}


func CreateComplaint(c *gin.Context) {
	var complaint models.Complaint
	if err := c.ShouldBindJSON(&complaint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&complaint).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, complaint)
}


func UpdateComplaint(c *gin.Context) {
	complaintID := c.Param("id")
	var complaint models.Complaint

	
	if err := database.DB.First(&complaint, complaintID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Complaint not found"})
		return
	}

	
	if err := c.ShouldBindJSON(&complaint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	if err := database.DB.Save(&complaint).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Complaint updated successfully", "complaint": complaint})
}


func DeleteComplaint(c *gin.Context) {
	complaintID := c.Param("id")
	var complaint models.Complaint

	
	if err := database.DB.First(&complaint, complaintID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Complaint not found"})
		return
	}


	if err := database.DB.Delete(&complaint).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Complaint deleted successfully"})
}
