package services

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
