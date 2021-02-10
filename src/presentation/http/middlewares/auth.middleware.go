package middlewares

import (
	jwtauth "main/src/infra/utils/jwt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// AuthMiddleware valid header token value
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		accessToken := strings.Split(authorizationHeader, " ")[1]

		claims := &jwtauth.Claims{}
		jwtSecretKey := []byte(os.Getenv("JWT_SECRET_KEY"))
		token, err := jwt.ParseWithClaims(accessToken, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecretKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		r.Header.Set("userID", claims.UserID)
		next.ServeHTTP(w, r)
	})
}
