package routes

import (
	"MentalHealthCare/controllers"

	"github.com/labstack/echo/v4"
)

func SetupRoutes(router *echo.Echo) {

	router.GET("/users", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUserByID)
	router.POST("/users", controllers.CreateUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	router.GET("/doctors", controllers.GetDoctors)
	router.GET("/doctors/:id", controllers.GetDoctorByID)
	router.POST("/doctors", controllers.CreateDoctor)
	router.PUT("/doctors/:id", controllers.UpdateDoctor)
	router.DELETE("/doctors/:id", controllers.DeleteDoctor)

	router.GET("/complaints", controllers.GetComplaints)
	router.GET("/complaints/:id", controllers.GetComplaintByID)
	router.POST("/complaints", controllers.CreateComplaint)
	router.PUT("/complaints/:id", controllers.UpdateComplaint)
	router.DELETE("/complaints/:id", controllers.DeleteComplaint)

	router.GET("/recommendations", controllers.GetRecommendations)
	router.GET("/recommendations/:id", controllers.GetRecommendationByID)
	router.POST("/recommendations", controllers.CreateRecommendation)
	router.PUT("/recommendations/:id", controllers.UpdateRecommendation)
	router.DELETE("/recommendations/:id", controllers.DeleteRecommendation)

	router.GET("/mental-health-info", controllers.GetMentalHealthInfo)
	router.GET("/mental-health-info/:id", controllers.GetMentalHealthInfoByID)
	router.POST("/mental-health-info", controllers.CreateMentalHealthInfo)
	router.PUT("/mental-health-info/:id", controllers.UpdateMentalHealthInfo)
	router.DELETE("/mental-health-info/:id", controllers.DeleteMentalHealthInfo)
}
