package main

import (
	"MentalHealthCare/database"
	"MentalHealthCare/routes"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	
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
