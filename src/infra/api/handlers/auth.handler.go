package handlers

import (
	"encoding/json"
	"io/ioutil"
	"main/src/application/controllers"
	"main/src/application/dtos"
	"net/http"
)

// PostAuthHandler handle post http request on auth route
func PostAuthHandler(resWriter http.ResponseWriter, req *http.Request) {
	var authCredentialsDto dtos.AuthCredentialsDto
	var authController controllers.AuthController

	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &authCredentialsDto)
	authResponseDto, err := authController.TryAuthenticate(authCredentialsDto)

	if err != nil {
		resWriter.WriteHeader(err.StatusCode)
		json.NewEncoder(resWriter).Encode(err)
	}

	resWriter.WriteHeader(http.StatusCreated)
	json.NewEncoder(resWriter).Encode(authResponseDto)
}
