package utils

import "github.com/dgrijalva/jwt-go"

// Claims is a JWT struct
type Claims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}
