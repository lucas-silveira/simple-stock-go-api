package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	. "main/src/application/controllers"
	. "main/src/application/dtos"
)

// GetUsers handle get http request on users route and delivery a list of users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header.Get("userID"))
	fmt.Fprintf(w, "Lista de usuários...")
}

// PostUsers handle post http request on users route and create a user
func PostUsers(w http.ResponseWriter, r *http.Request) {
	var createUserDto CreateUserDto
	var userController UserController

	json.NewDecoder(r.Body).Decode(&createUserDto)
	err := userController.CreateAnUser(createUserDto)

	if err != nil {
		w.WriteHeader(err.StatusCode)
		json.NewEncoder(w).Encode(err)
	}

	w.WriteHeader(http.StatusCreated)
}
