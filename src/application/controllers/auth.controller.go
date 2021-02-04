package controllers

import (
	"fmt"
	"main/src/application/dtos"
)

// AuthController ...
type AuthController struct{}

// TryAuthenticate ...
func (authController AuthController) TryAuthenticate(authCredentialsDto dtos.AuthCredentialsDto) {
	fmt.Println(authCredentialsDto)
}
