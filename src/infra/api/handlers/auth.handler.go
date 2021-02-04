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
	var authRequestDto dtos.AuthRequestDto
	var authController controllers.AuthController

	body, _ := ioutil.ReadAll(req.Body)
	json.Unmarshal(body, &authRequestDto)
	authController.TryAuthenticate(authRequestDto)

	resWriter.WriteHeader(http.StatusCreated)
}
