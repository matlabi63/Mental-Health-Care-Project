package database

import (
	"MentalHealthCare/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
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
    fmt.Printf("DB_USER: %s\n", user)
    fmt.Printf("DB_PASSWORD: %s\n", password)
    fmt.Printf("DB_HOST: %s\n", host)
    fmt.Printf("DB_PORT: %s\n", port)
    fmt.Printf("DB_NAME: %s\n", dbname)

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        user, password, host, port, dbname)

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect to database: " + err.Error())
    }

    // Migrate the schema
    DB.AutoMigrate(&models.User{}, &models.Doctor{}, &models.Complaint{}, &models.Recommendation{}, &models.Information{}, &models.Admin{})
}