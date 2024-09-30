package main

import (
	"MentalHealthCare/database"
	"MentalHealthCare/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.InitDatabase()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.SetupRoutes(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "2313"
	}

	e.Logger.Fatal(e.Start(":" + port))
}

// package main

// import (
// 	"MentalHealthCare/database"
// 	"MentalHealthCare/routes"

// 	"github.com/labstack/echo/v4"
// )

// func main() {
// 	database.InitDB()

// 	e := echo.New()

// 	routes.SetupRoutes(e)

// 	e.Logger.Fatal(e.Start("1323"))
// }