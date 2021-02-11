package controllers_test

import (
	"main/src/application/controllers"
	"main/src/application/dtos"
	"main/src/infra/data/repositories"
	"main/src/infra/errors"
	"main/src/infra/utils/encryptor"
	"net/http"
	"reflect"
	"testing"
)

func TestCreateAnUser(t *testing.T) {
	t.Parallel()

	msgIndex := "test: %s failed - expected (%v) <> found (%v)."

	fakeRepo := repositories.NewFakeUser()
	encryptor := &encryptor.BcryptEncryptor{}
	userController := &controllers.UserController{fakeRepo, encryptor}

	tests := []struct {
		name    string
		request dtos.CreateUserDto
		expect  *errors.Http
	}{
		{
			name: "Should be able to create a new user successfully",
			request: dtos.CreateUserDto{
				Name:            "John",
				Email:           "john@email.com",
				Password:        "123456",
				ConfirmPassword: "123456",
			},
			expect: nil,
		},
		{
			name: "Should not be able to create a new user if passwords doesn't match",
			request: dtos.CreateUserDto{
				Name:            "John",
				Email:           "john@email.com",
				Password:        "123456",
				ConfirmPassword: "1234567",
			},
			expect: errors.NewHttpError(
				http.StatusBadRequest,
				"Existem alguns campos inválidos.",
				[]string{"ConfirmPassword deve ser igual a Password"},
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := userController.CreateAnUser(test.request)

			if !reflect.DeepEqual(err, test.expect) {
				t.Errorf(msgIndex, test.name, test.expect, err)
			}
		})
	}
}
