package middlewares

import (
	"fmt"
	"net/http"
)

// AuthMiddleware valid header token value
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Passou aqui")

		next.ServeHTTP(w, r)
	})
}
