// package utils

// import (
// 	"crypto/rand"
// 	"encoding/base64"
// 	"log"
// 	"os"
// 	"time"

// 	"github.com/spf13/viper"
// 	"golang.org/x/crypto/bcrypt"
// )

// // GetConfig retrieves a configuration value from the .env file.
// func GetConfig(key string) string {
// 	viper.SetConfigFile(".env")
// 	if err := viper.ReadInConfig(); err != nil {
// 		log.Fatalf("error reading config file: %s", err)
// 	}
// 	return viper.GetString(key)
// }

// // GenerateRandomString generates a random string of the specified length.
// func GenerateRandomString(length int) (string, error) {
// 	bytes := make([]byte, length)
// 	_, err := rand.Read(bytes)
// 	if err != nil {
// 		return "", err
// 	}
// 	return base64.RawStdEncoding.EncodeToString(bytes)[:length], nil
// }

// // CurrentTimestamp returns the current timestamp.
// func CurrentTimestamp() int64 {
// 	return time.Now().Unix()
// }

// // CheckError logs the error and exits if it's not nil.
// func CheckError(err error) {
// 	if err != nil {
// 		log.Fatalf("an error occurred: %s", err)
// 	}
// }

// // IsEnvProduction checks if the application is running in production.
// func IsEnvProduction() bool {
// 	return os.Getenv("ENV") == "production"
// }

// func HashPassword(password string) (string, error) {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	return string(bytes), err
// }

// func CheckPasswordHash(password, hash string) bool {
// 	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
// 	return err == nil
// }

//

package utils

import (
	"log"

	"github.com/spf13/viper"
)

func GetConfig(key string) string {
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error while reading config file %s\n", err)
	}

	return viper.GetString(key)
}

