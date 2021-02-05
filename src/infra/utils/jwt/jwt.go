package jwtauth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateToken generate a new jwt token
func GenerateToken(userID string) (string, error) {
	expirationTime := time.Now().Add(2 * 24 * time.Hour)
	claims := NewClaims(userID, expirationTime)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(jwtSecretKey)

	return tokenString, err
}
