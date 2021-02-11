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
	"main/src/infra/validator"
	"net/http"
)

// UserController ...
type UserController struct {
	UserRepository user.IRepository
	Encryptor      domain.IEncryptor
}

// NewUserController is a factory that return a new UserController object
func NewUserController() UserController {
	userRepository := repositories.NewUserRepository(drivers.DBConnection)
	encryptor := &encryptor.BcryptEncryptor{}
	return UserController{userRepository, encryptor}
}

// CreateAnUser ...
func (userController UserController) CreateAnUser(createUserDto CreateUserDto) *errors.Http {
	var err error
	errorsMessage, err := validator.Validate(createUserDto)

	if err != nil {
		return errors.NewHttpError(
			http.StatusBadRequest,
			"Existem alguns campos inválidos.",
			errorsMessage,
		)
	}

	userExists, err := userController.UserRepository.FindByEmail(createUserDto.Email)

	if err != nil {
		fmt.Println(err)
		return errors.NewHttpError(
			http.StatusInternalServerError,
			"Internal server error.",
			[]string{},
		)
	}

	if *userExists != (user.User{}) {
		return errors.NewHttpError(
			http.StatusBadRequest,
			"The user already exists.",
			[]string{},
		)
	}

	passwordHashed, err := userController.Encryptor.CreateHash(createUserDto.Password)

	if err != nil {
		fmt.Println(err)
		return errors.NewHttpError(
			http.StatusInternalServerError,
			"Internal server error.",
			[]string{},
		)
	}

	newUser := user.NewUser(
		createUserDto.Name,
		createUserDto.Email,
		passwordHashed,
	)

	err = userController.UserRepository.Save(newUser)

	if err != nil {
		fmt.Println(err)
		return errors.NewHttpError(
			http.StatusInternalServerError,
			"Internal server error.",
			[]string{},
		)
	}

	return nil
}
