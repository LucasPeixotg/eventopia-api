package common

import (
	"net/http"
	"time"
)

type ApiError struct {
	Error string `json:"error"`
}
type ApiFunc func(http.ResponseWriter, *http.Request) (int, error)
type ApiFuncWithAuth func(http.ResponseWriter, *http.Request, Account) (int, error)

/*
Tables
*/

type Account struct {
	ID   int
	CPF  string
	Name string
}

type Event struct {
	ID          int       `json:"id"`
	OwnerID     int       `json:"owner_id"`
	Name        string    `json:"name"`
	DateTime    time.Time `json:"datetime"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
}

/*
Requests
*/

type CreateAccountRequest struct {
	Cpf      string `json:"cpf"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Cpf      string `json:"cpf"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type CreateEventRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"datetime"`
}
