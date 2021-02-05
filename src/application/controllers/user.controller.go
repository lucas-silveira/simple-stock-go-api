package controllers

import (
	"fmt"
	. "main/src/application/dtos"
	. "main/src/domain/user"
	"main/src/infra/errors"
	"main/src/infra/validator"
	"net/http"
)

// UserController ...
type UserController struct{}

// CreateAnUser ...
func (userController UserController) CreateAnUser(createUserDto CreateUserDto) *errors.Http {
	errorsMessage, err := validator.Validate(createUserDto)

	if err != nil {
		return errors.NewHttpError(
			http.StatusBadRequest,
			"Existem alguns campos inválidos.",
			errorsMessage,
		)
	}

	newUser := NewUser(
		createUserDto.Name,
		createUserDto.Email,
		createUserDto.Password,
	)
	fmt.Println(newUser)

	return nil
}
