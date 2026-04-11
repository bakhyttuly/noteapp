package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JwtKey = []byte("my_secret_key")

func GenerateToken(userID uint, username string) (string, error) {
	claims := jwt.MapClaims{
		"id":       userID,
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
