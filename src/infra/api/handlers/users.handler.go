package handlers

import (
	"fmt"
	"net/http"
)

// GetAllUsers handle get http request on users route and delivery a list of users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Lista de usuários...")
}
