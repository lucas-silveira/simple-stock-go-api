package jwtauth

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims is a JWT object
type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// NewClaims is a factory function that generate a new Claims
func NewClaims(userID string, expirationTime time.Time) Claims {
	return Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
}
