package controllers

import (
	"fmt"
	. "main/src/application/dtos"
	. "main/src/domain/user"
	"main/src/infra/data/drivers"
	"main/src/infra/data/repositories"
	"main/src/infra/errors"
	"main/src/infra/validator"
	"net/http"
)

// UserController ...
type UserController struct{}

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

	var userRepository IRepository

	userRepository = repositories.NewUserRepository(drivers.DBConnection)

	newUser := NewUser(
		createUserDto.Name,
		createUserDto.Email,
		createUserDto.Password,
	)

	userExists, err := userRepository.FindByEmail(newUser.Email)

	if err != nil {
		fmt.Println(err)
		return errors.NewHttpError(
			http.StatusInternalServerError,
			"Internal server error.",
			[]string{},
		)
	}

	if userExists != nil {
		return errors.NewHttpError(
			http.StatusBadRequest,
			"The user already exists.",
			[]string{},
		)
	}

	err = userRepository.Save(newUser)

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
