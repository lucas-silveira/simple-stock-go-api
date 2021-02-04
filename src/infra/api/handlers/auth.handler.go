package handlers

import (
	"encoding/json"
	"main/src/application/controllers"
	"main/src/application/dtos"
	"net/http"
)

// PostAuth handle post http request on auth route
func PostAuth(w http.ResponseWriter, r *http.Request) {
	var authCredentialsDto dtos.AuthCredentialsDto
	var authController controllers.AuthController

	json.NewDecoder(r.Body).Decode(&authCredentialsDto)
	authResponseDto, err := authController.TryAuthenticate(authCredentialsDto)

	if err != nil {
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(err)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(authResponseDto)
}
