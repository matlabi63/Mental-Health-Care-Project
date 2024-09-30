package database

import (
	"MentalHealthCare/models"
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
    user := os.Getenv("DB_USER")
    password := os.Getenv("DB_PASSWORD")
    host := os.Getenv("DB_HOST")
    port := os.Getenv("DB_PORT")
    dbname := os.Getenv("DB_NAME")

    // Debug output
    // fmt.Printf("DB_USER: %s\n", user)
    // fmt.Printf("DB_PASSWORD: %s\n", password)
    // fmt.Printf("DB_HOST: %s\n", host)
    // fmt.Printf("DB_PORT: %s\n", port)
    // fmt.Printf("DB_NAME: %s\n", dbname)

    var dsn string = fmt.Sprintf(
	    "user=%s password=%s host=%s port=%s dbname=%s",
		user,
		password,
		host,
		port,
		dbname,
	)

    pgConfig := postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})

    var err error

	DB, err = gorm.Open(pgConfig, &gorm.Config{})

   

    if err != nil {
        panic("failed to connect to database: " + err.Error())
    }

    // Migrate the schema
    DB.AutoMigrate(&models.User{}, &models.Doctor{}, &models.Complaint{}, &models.Recommendation{}, &models.Information{}, &models.Admin{})
}
