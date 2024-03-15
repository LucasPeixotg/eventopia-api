package db

import (
	"sync"
	c "vistaverse/src/common"
)

type Storage interface {
	CreateAccount(*c.CreateAccountRequest) (*c.Account, error)
	Login(*c.LoginRequest) (*c.Account, error)

	CreateEvent(req *c.CreateEventRequest, user_id int) (*c.Event, error)
}

var storage Storage

var lock = &sync.Mutex{}

// implements singleton pattern to always get same storage
func GetStorage() Storage {
	if storage == nil {
		lock.Lock()
		defer lock.Unlock()
		if storage == nil {
			// creates storage instance
			storage = newPostgressStoreOrFatal()
		}
	}
	return storage
}
