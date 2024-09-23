package utils

import (
	"errors"
	"time"

	middlewares "MentalHealthCare/middlewares"
	"MentalHealthCare/models"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateJWT creates a new JWT token for a user.
func GenerateJWT(userID int, jwtOptions models.JWTOptions) (string, error) {
	expirationTime := time.Now().Add(time.Duration(jwtOptions.ExpiresDuration) * time.Hour)

	claims := &middlewares.JWTCustomClaims{
		ID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtOptions.SecretKey))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

// ValidateJWT checks the validity of the JWT token.
func ValidateJWT(tokenString string, jwtOptions models.JWTOptions) (*middlewares.JWTCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &middlewares.JWTCustomClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtOptions.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*middlewares.JWTCustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// ExtractToken retrieves the token from the Authorization header.
func ExtractToken(authorizationHeader string) (string, error) {
	if authorizationHeader == "" {
		return "", errors.New("authorization header is empty")
	}

	token := authorizationHeader[len("Bearer "):]
	if len(token) == 0 {
		return "", errors.New("token is missing")
	}

	return token, nil
}
