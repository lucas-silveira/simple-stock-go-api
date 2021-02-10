package controllers

import (
	"fmt"
	. "main/src/application/dtos"
	"main/src/domain"
	"main/src/domain/user"
	"main/src/infra/data/drivers"
	"main/src/infra/data/repositories"
	"main/src/infra/errors"
	"main/src/infra/utils/encryptor"
	jwtauth "main/src/infra/utils/jwt"
	"net/http"
	"strconv"
)

// AuthController ...
type AuthController struct {
	userRepository user.IRepository
	encryptor      domain.IEncryptor
}

// NewAuthController is a factory that return a new AuthController object
func NewAuthController() AuthController {
	userRepository := repositories.NewUserRepository(drivers.DBConnection)
	encryptor := &encryptor.BcryptEncryptor{}
	return AuthController{userRepository, encryptor}
}

// TryAuthenticate ...
func (authController AuthController) TryAuthenticate(authCredentialsDto AuthCredentialsDto) (AuthResponseDto, *errors.Http) {
	userExists, err := authController.userRepository.FindByEmail(authCredentialsDto.Email)

	if err != nil {
		fmt.Println(err)
		return AuthResponseDto{}, errors.NewHttpError(
			http.StatusInternalServerError,
			"Internal server error.",
			[]string{},
		)
	}

	if *userExists == (user.User{}) {
		return AuthResponseDto{}, errors.NewHttpError(
			http.StatusNotFound,
			"User not found.",
			[]string{},
		)
	}

	passwordIsValid := authController.encryptor.Validate(userExists.Password, authCredentialsDto.Password)

	if !passwordIsValid {
		return AuthResponseDto{}, errors.NewHttpError(
			http.StatusUnauthorized,
			"Email or password invalid.",
			[]string{},
		)
	}

	accessToken, err := jwtauth.GenerateToken(strconv.Itoa(userExists.ID))

	if err != nil {
		return AuthResponseDto{},
			errors.NewHttpError(
				http.StatusInternalServerError,
				"",
				[]string{},
			)
	}
	authResponseDto := NewAuthResponseDto(userExists.ID, accessToken)

	return authResponseDto, nil
}
