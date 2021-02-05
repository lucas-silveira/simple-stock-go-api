package controllers

import (
	"fmt"
	. "main/src/application/dtos"
	. "main/src/domain/user"
	"main/src/infra/errors"
)

// UserController ...
type UserController struct{}

// CreateAnUser ...
func (userController UserController) CreateAnUser(createUserDto CreateUserDto) *errors.Http {
	newUser := NewUser(createUserDto.Name, createUserDto.Email, createUserDto.Password)
	fmt.Println(newUser)

	return nil
}
