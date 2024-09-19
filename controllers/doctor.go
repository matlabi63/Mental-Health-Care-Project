package controllers

import (
	"MentalHealthCare/database"
	"MentalHealthCare/models"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetDoctors(c *gin.Context) {
	var doctors []models.Doctor
	if err := database.DB.Find(&doctors).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, doctors)
}


func GetDoctorByID(c *gin.Context) {
	doctorID := c.Param("id")
	var doctor models.Doctor

	if err := database.DB.First(&doctor, doctorID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}
	c.JSON(http.StatusOK, doctor)
}


func CreateDoctor(c *gin.Context) {
	var doctor models.Doctor
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	var user models.User
	if err := database.DB.First(&user, doctor.UserID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	if err := database.DB.Create(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, doctor)
}


func UpdateDoctor(c *gin.Context) {
	doctorID := c.Param("id")
	var doctor models.Doctor


	if err := database.DB.First(&doctor, doctorID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}

	
	if err := c.ShouldBindJSON(&doctor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	
	if doctor.UserID != 0 {
		var user models.User
		if err := database.DB.First(&user, doctor.UserID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
			return
		}
	}

	
	if err := database.DB.Save(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Doctor updated successfully", "doctor": doctor})
}


func DeleteDoctor(c *gin.Context) {
	doctorID := c.Param("id")
	var doctor models.Doctor

	
	if err := database.DB.First(&doctor, doctorID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Doctor not found"})
		return
	}

	
	if err := database.DB.Delete(&doctor).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Doctor deleted successfully"})
}
