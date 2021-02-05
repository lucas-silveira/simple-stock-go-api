package controllers

import (
	. "main/src/application/dtos"
	"main/src/infra/errors"
	"main/src/infra/utils"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// AuthController ...
type AuthController struct{}

// TryAuthenticate ...
func (authController AuthController) TryAuthenticate(authCredentialsDto AuthCredentialsDto) (AuthResponseDto, *errors.Http) {
	expirationTime := time.Now().Add(2 * 24 * time.Hour)
	userID := "1"
	claims := &utils.Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtSecretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(jwtSecretKey)

	if err != nil {
		return AuthResponseDto{}, &errors.Http{
			StatusCode: http.StatusInternalServerError,
			Message:    "",
		}
	}

	userIDParsed, _ := strconv.Atoi(userID)
	authResponseDto := NewAuthResponseDto(userIDParsed, tokenString)

	return authResponseDto, nil
}
