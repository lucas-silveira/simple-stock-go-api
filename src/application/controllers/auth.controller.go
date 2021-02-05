package controllers

import (
	. "main/src/application/dtos"
	"main/src/infra/errors"
	jwtauth "main/src/infra/utils/jwt"
	"net/http"
	"strconv"
)

// AuthController ...
type AuthController struct{}

// TryAuthenticate ...
func (authController AuthController) TryAuthenticate(authCredentialsDto AuthCredentialsDto) (AuthResponseDto, *errors.Http) {
	userID := "1"
	accessToken, err := jwtauth.GenerateToken(userID)

	if err != nil {
		return AuthResponseDto{}, &errors.Http{
			StatusCode: http.StatusInternalServerError,
			Message:    "",
			Errors:     []string{},
		}
	}

	userIDParsed, _ := strconv.Atoi(userID)
	authResponseDto := NewAuthResponseDto(userIDParsed, accessToken)

	return authResponseDto, nil
}
