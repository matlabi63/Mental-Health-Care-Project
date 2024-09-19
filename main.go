package main

import (
	"MentalHealthCare/database"
	"MentalHealthCare/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    // Load .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // Initialize the database
    database.InitDatabase()

    // Create a new Gin router
    router := gin.Default()

    // Setup routes
    routes.SetupRoutes(router)

    // Start the server
    router.Run(":8080")
}