package auth

import (
	"fmt"
	"log"
	"net/http"
	"vistaverse/src/common"

	"github.com/golang-jwt/jwt/v5"
)

func WithLogin(f common.ApiFuncWithAuth) common.ApiFunc {
	return func(w http.ResponseWriter, r *http.Request) (int, error) {
		tokenString := r.Header.Get("Authorization")

		// missing authorization header
		if tokenString == "" {
			return http.StatusUnauthorized, fmt.Errorf("unauthorized")
		}

		tokenString = tokenString[len("Bearer "):]

		// invalid token
		token, err := verifyToken(tokenString)
		if err != nil {
			return http.StatusUnauthorized, fmt.Errorf("unauthorized")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			log.Printf("ERROR: could not parse claims\n")
			return http.StatusUnauthorized, fmt.Errorf("unauthorized")
		}

		acc := common.Account{
			ID:   int(claims["ID"].(float64)),
			CPF:  claims["CPF"].(string),
			Name: claims["Name"].(string),
		}
		return f(w, r, acc)
	}
}
