package controllers

import (
	"fmt"
	"main/src/application/dtos"
)

// AuthController ...
type AuthController struct{}

// TryAuthenticate ...
func (authController AuthController) TryAuthenticate(authRequestDto dtos.AuthRequestDto) {
	fmt.Println(authRequestDto)
}
