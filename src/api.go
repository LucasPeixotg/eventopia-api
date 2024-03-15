package src

import (
	"log"
	"net/http"
	"vistaverse/src/auth"
	"vistaverse/src/event"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func RunAPI(address string) {
	router := mux.NewRouter()

	auth.SetAuthRoutes(router.PathPrefix("/auth").Subrouter())
	event.SetEventRoutes(router.PathPrefix("/event").Subrouter())

	handler := cors.Default().Handler(router)

	log.Printf("server started. listening at %s\n", address)
	// continues to serve indefinetly and prints the error on fatal
	log.Fatal(http.ListenAndServe(address, handler))
}
