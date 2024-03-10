package src

import (
	"log"
	"net/http"
	"vistaverse/src/auth"

	"github.com/gorilla/mux"
)

func RunAPI(address string) {
	router := mux.NewRouter()

	auth.SetAuthRoutes(router.PathPrefix("/auth").Subrouter())

	log.Printf("Server started... Listening at: %s\n", address)
	// continues to serve indefinetly and prints the error on fatal
	log.Fatal(http.ListenAndServe(address, router))
}
