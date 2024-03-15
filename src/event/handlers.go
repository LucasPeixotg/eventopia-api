package event

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/LucasPeixotg/eventopia-api/src/auth"
	c "github.com/LucasPeixotg/eventopia-api/src/common"
	"github.com/LucasPeixotg/eventopia-api/src/db"

	"github.com/gorilla/mux"
)

var store = db.GetStorage()

func SetEventRoutes(router *mux.Router) {
	router.HandleFunc("", c.MakeHTTPHandlerFunc(auth.WithLogin(handleBaseURL)))
}

func handleBaseURL(w http.ResponseWriter, r *http.Request, acc c.Account) (int, error) {
	switch r.Method {
	case "GET":
		return handleGetEvents(w, r, acc)
	case "POST":
		return handleCreateEvent(w, r, acc)
	default:
		return http.StatusMethodNotAllowed, fmt.Errorf("method not allowed %s", r.Method)
	}
}

func handleCreateEvent(w http.ResponseWriter, r *http.Request, acc c.Account) (int, error) {
	req := &c.CreateEventRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("ERROR: could not decode JSON %s\n", err.Error())
		return http.StatusBadRequest, fmt.Errorf("bad request")
	}

	event, err := store.CreateEvent(req, acc.ID)
	if err != nil {
		log.Printf("ERROR: could not create event %s\n", err)
		return http.StatusInternalServerError, fmt.Errorf("could not create event")
	}

	if err := c.WriteJSON(w, http.StatusOK, event); err != nil {
		log.Printf("ERROR: could not write JSON %s\n", err.Error())
		return http.StatusInternalServerError, fmt.Errorf("something unexpected happened")
	}

	return http.StatusOK, nil
}

func handleGetEvents(w http.ResponseWriter, r *http.Request, acc c.Account) (int, error) {
	return 0, nil
}
