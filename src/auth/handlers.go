package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	c "github.com/LucasPeixotg/eventopia-api/src/common"
	"github.com/LucasPeixotg/eventopia-api/src/db"
	"github.com/gorilla/mux"
)

var store = db.GetStorage()

func SetAuthRoutes(router *mux.Router) {
	router.HandleFunc("/register", c.MakeHTTPHandlerFunc(handleCreateAccount))
	router.HandleFunc("/login", c.MakeHTTPHandlerFunc(handleLogin))
}

func handleCreateAccount(w http.ResponseWriter, r *http.Request) (int, error) {
	if r.Method != "POST" {
		return http.StatusMethodNotAllowed, fmt.Errorf("method not allowed %s", r.Method)
	}
	req := &c.CreateAccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return http.StatusBadRequest, fmt.Errorf("bad request")
	}

	acc, err := store.CreateAccount(req)
	if err != nil {
		log.Printf("ERROR: (could not write create account) %s\n", err.Error())
		return http.StatusBadRequest, fmt.Errorf("could not create account")
	}

	// auto login created account
	token, err := signToken(acc)
	if err != nil {
		log.Printf("ERROR: (could not sign token) %s\n", err.Error())
		return http.StatusForbidden, fmt.Errorf("could not login")
	}

	if err := c.WriteJSON(w, http.StatusOK, tokenResponse{token, acc.Name}); err != nil {
		return http.StatusForbidden, fmt.Errorf("could not login")
	}

	return http.StatusOK, nil
}

func handleLogin(w http.ResponseWriter, r *http.Request) (int, error) {
	if r.Method != "POST" {
		return http.StatusMethodNotAllowed, fmt.Errorf("method not allowed %s", r.Method)
	}

	req := &c.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return http.StatusBadRequest, fmt.Errorf("bad request")
	}

	acc, err := store.Login(req)
	if err != nil {
		return http.StatusForbidden, fmt.Errorf("could not login")
	}

	token, err := signToken(acc)
	if err != nil {
		log.Printf("ERROR: (could not sign token) %s\n", err.Error())
		return http.StatusForbidden, fmt.Errorf("could not login")
	}

	if err := c.WriteJSON(w, http.StatusOK, tokenResponse{token, acc.Name}); err != nil {
		return http.StatusForbidden, fmt.Errorf("could not login")
	}

	return http.StatusOK, nil
}
