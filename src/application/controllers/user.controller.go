package controllers

import (
	"fmt"
	. "main/src/application/dtos"
	"main/src/domain"
	"main/src/domain/user"
	. "main/src/domain/user"
	"main/src/infra/data/drivers"
	"main/src/infra/data/repositories"
	"main/src/infra/errors"
	"main/src/infra/utils/encryptor"
	"main/src/infra/validator"
	"net/http"
)

// UserController ...
type UserController struct {
	userRepository user.IRepository
	encryptor      domain.IEncryptor
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

	userExists, err := userController.userRepository.FindByEmail(createUserDto.Email)

	if err != nil {
		fmt.Println(err)
		return errors.NewHttpError(
			http.StatusInternalServerError,
			"Internal server error.",
			[]string{},
		)
	}

	if *userExists != (User{}) {
		return errors.NewHttpError(
			http.StatusBadRequest,
			"The user already exists.",
			[]string{},
		)
	}

	passwordHashed, err := userController.encryptor.CreateHash(createUserDto.Password)

	if err != nil {
		fmt.Println(err)
		return errors.NewHttpError(
			http.StatusInternalServerError,
			"Internal server error.",
			[]string{},
		)
	}

	newUser := NewUser(
		createUserDto.Name,
		createUserDto.Email,
		passwordHashed,
	)

	err = userController.userRepository.Save(newUser)

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
