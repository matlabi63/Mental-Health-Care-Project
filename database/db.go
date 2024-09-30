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

// package database

// import (
// 	"MentalHealthCare/models"
// 	"MentalHealthCare/utils"
// 	"fmt"
// 	"log"

// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// var (
// 	DB_USERNAME string = utils.GetConfig("DB_USERNAME")
// 	DB_PASSWORD string = utils.GetConfig("DB_PASSWORD")
// 	DB_NAME     string = utils.GetConfig("DB_NAME")
// 	DB_HOST     string = utils.GetConfig("DB_HOST")
// 	DB_PORT     string = utils.GetConfig("DB_PORT")
// )

// func InitDB() {
// 	var err error

// 	var dsn string = fmt.Sprintf(
// 		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
// 		DB_USERNAME,
// 		DB_PASSWORD,
// 		DB_HOST,
// 		DB_PORT,
// 		DB_NAME,
// 	)

// 	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatalf("error when connecting to the database: %s\n", err)
// 	}

// 	log.Println("connected to the database")

// 	Migrate()
// }

// func Migrate() {
// 	err := DB.AutoMigrate(&models.Admin{}, &models.Complaint{}, &models.Doctor{}, &models.Information{}, &models.Recommendation{}, &models.User{})

// 	if err != nil {
// 		log.Fatalf("error when migration: %s\n", err)
// 	}

// 	log.Println("migration successful")
// }