package routes

import (
	"MentalHealthCare/controllers"
	"MentalHealthCare/middlewares"
	"MentalHealthCare/models"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo) {
	// Initialize logger middleware
	loggerConfig := middlewares.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	loggerMiddleware := loggerConfig.Init()
	e.Use(loggerMiddleware)

	// Initialize authentication middleware
	jwtConfig := middlewares.JWTConfig{
		SecretKey: os.Getenv("JWT_SECRET_KEY"),
	}

	authMiddlewareConfig := jwtConfig.Init()

	jwtOptions := models.JWTOptions{
		SecretKey:       os.Getenv("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	// User routes
	userController := controllers.InitUserController(jwtOptions)
	userRoutes := e.Group("/api/v1/users")

	userRoutes.POST("/register", userController.Register)
	userRoutes.POST("/login", userController.Login)
	userRoutes.GET(
		"/info",
		userController.GetUserInfo,
		echojwt.WithConfig(authMiddlewareConfig),
		middlewares.VerifyToken,
	)

// Admin routes
adminController := controllers.InitAdminController()
adminRoutes := e.Group("/api/v1/admin", echojwt.WithConfig(authMiddlewareConfig), middlewares.VerifyToken)

adminRoutes.GET("", adminController.GetAll)
adminRoutes.GET("/:id", adminController.GetByID)
adminRoutes.POST("", adminController.Create)
adminRoutes.PUT("/:id", adminController.Update)
adminRoutes.DELETE("/:id", adminController.Delete)

	// Doctor routes
	doctorController := controllers.InitDoctorController()
	doctorRoutes := e.Group("/api/v1/doctors", echojwt.WithConfig(authMiddlewareConfig), middlewares.VerifyToken)

	doctorRoutes.GET("", doctorController.GetAll)
	doctorRoutes.GET("/:id", doctorController.GetByID)
	doctorRoutes.POST("", doctorController.Create)
	doctorRoutes.PUT("/:id", doctorController.Update)
	doctorRoutes.DELETE("/:id", doctorController.Delete)

	// Complaint routes
	complaintController := controllers.InitComplaintController()
	complaintRoutes := e.Group("/api/v1/complaints", echojwt.WithConfig(authMiddlewareConfig), middlewares.VerifyToken)

	complaintRoutes.GET("", complaintController.GetAll)
	complaintRoutes.GET("/:id", complaintController.GetByID)
	complaintRoutes.POST("", complaintController.Create)
	complaintRoutes.PUT("/:id", complaintController.Update)
	complaintRoutes.DELETE("/:id", complaintController.Delete)

	// Recommendation routes
	recommendationController := controllers.InitRecommendationController()
	recommendationRoutes := e.Group("/api/v1/recommendations", echojwt.WithConfig(authMiddlewareConfig), middlewares.VerifyToken)

	recommendationRoutes.GET("", recommendationController.GetAll)
	recommendationRoutes.GET("/:id", recommendationController.GetByID)
	recommendationRoutes.POST("", recommendationController.Create)
	recommendationRoutes.PUT("/:id", recommendationController.Update)
	recommendationRoutes.DELETE("/:id", recommendationController.Delete)

	// information routes
	InformationController := controllers.InitInformationController()
	InformationRoutes := e.Group("/api/v1/mental-health-info", echojwt.WithConfig(authMiddlewareConfig), middlewares.VerifyToken)

	InformationRoutes.GET("", InformationController.GetAll)
	InformationRoutes.GET("/:id", InformationController.GetByID)
	InformationRoutes.POST("", InformationController.Create)
	InformationRoutes.PUT("/:id", InformationController.Update)
	InformationRoutes.DELETE("/:id", InformationController.Delete)
}