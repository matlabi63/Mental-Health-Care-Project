<!-- routes

package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Users Endpoints
    router.GET("/users", getUsers)
    router.GET("/users/:id", getUserByID)
    router.POST("/users", createUser)
    router.PUT("/users/:id", updateUser)
    router.DELETE("/users/:id", deleteUser)

    // Doctors Endpoints
    router.GET("/doctors", getDoctors)
    router.GET("/doctors/:id", getDoctorByID)
    router.POST("/doctors", createDoctor)
    router.PUT("/doctors/:id", updateDoctor)
    router.DELETE("/doctors/:id", deleteDoctor)

    // Patients Endpoints
    router.GET("/patients", getPatients)
    router.GET("/patients/:id", getPatientByID)
    router.POST("/patients", createPatient)
    router.PUT("/patients/:id", updatePatient)
    router.DELETE("/patients/:id", deletePatient)

    // Complaints Endpoints
    router.GET("/complaints", getComplaints)
    router.GET("/complaints/:id", getComplaintByID)
    router.POST("/complaints", createComplaint)
    router.PUT("/complaints/:id", updateComplaint)
    router.DELETE("/complaints/:id", deleteComplaint)

    // Recommendations Endpoints
    router.GET("/recommendations", getRecommendations)
    router.GET("/recommendations/:id", getRecommendationByID)
    router.POST("/recommendations", createRecommendation)
    router.PUT("/recommendations/:id", updateRecommendation)
    router.DELETE("/recommendations/:id", deleteRecommendation)

    // Informations Endpoints (Mental Health Articles)
    router.GET("/informations", getInformations)
    router.GET("/informations/:id", getInformationByID)
    router.POST("/informations", createInformation)
    router.PUT("/informations/:id", updateInformation)
    router.DELETE("/informations/:id", deleteInformation)

    // Admins Endpoints
    router.GET("/admins", getAdmins)
    router.GET("/admins/:id", getAdminByID)
    router.POST("/admins", createAdmin)
    router.PUT("/admins/:id", updateAdmin)
    router.DELETE("/admins/:id", deleteAdmin)

    router.Run(":8080")
}

func getUsers(c *gin.Context) {}
func getUserByID(c *gin.Context) {}
func createUser(c *gin.Context) {}
func updateUser(c *gin.Context) {}
func deleteUser(c *gin.Context) {}

func getDoctors(c *gin.Context) {}
func getDoctorByID(c *gin.Context) {}
func createDoctor(c *gin.Context) {}
func updateDoctor(c *gin.Context) {}
func deleteDoctor(c *gin.Context) {}

func getPatients(c *gin.Context) {}
func getPatientByID(c *gin.Context) {}
func createPatient(c *gin.Context) {}
func updatePatient(c *gin.Context) {}
func deletePatient(c *gin.Context) {}

func getComplaints(c *gin.Context) {}
func getComplaintByID(c *gin.Context) {}
func createComplaint(c *gin.Context) {}
func updateComplaint(c *gin.Context) {}
func deleteComplaint(c *gin.Context) {}

func getRecommendations(c *gin.Context) {}
func getRecommendationByID(c *gin.Context) {}
func createRecommendation(c *gin.Context) {}
func updateRecommendation(c *gin.Context) {}
func deleteRecommendation(c *gin.Context) {}

func getInformations(c *gin.Context) {}
func getInformationByID(c *gin.Context) {}
func createInformation(c *gin.Context) {}
func updateInformation(c *gin.Context) {}
func deleteInformation(c *gin.Context) {}

func getAdmins(c *gin.Context) {}
func getAdminByID(c *gin.Context) {}
func createAdmin(c *gin.Context) {}
func updateAdmin(c *gin.Context) {}
func deleteAdmin(c *gin.Context) {}
 -->
