package controllers

import (
	"main/src/application/dtos"
	"main/src/infra/errors"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AuthController ...
type AuthController struct{}

type Claims struct {
	userID int
	jwt.StandardClaims
}

// TryAuthenticate ...
func (authController AuthController) TryAuthenticate(authCredentialsDto dtos.AuthCredentialsDto) (*dtos.AuthResponseDto, *errors.Http) {
	expirationTime := time.Now().Add(2 * 24 * time.Hour)
	userID := 1
	claims := &Claims{
		userID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("my_secret_key"))

	if err != nil {
		return nil, &errors.Http{
			StatusCode: http.StatusInternalServerError,
			Message:    "",
		}
	}

	authResponseDto := dtos.NewAuthResponseDto(userID, tokenString)
	return authResponseDto, nil
}
